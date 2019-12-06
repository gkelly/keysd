package service

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/keys-pub/keys"
	"github.com/keys-pub/keys/keyring"
	"github.com/keys-pub/keys/saltpack"
	"github.com/keys-pub/keysd/http/client"
	"github.com/mercari/go-grpc-interceptor/panichandler"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type protoService struct {
	*service
}

func newProtoService(cfg *Config, build Build, auth *auth) (*protoService, error) {
	srv := newService(cfg, build, auth)
	if err := srv.Open(); err != nil {
		return nil, err
	}
	p := &protoService{srv}
	return p, nil
}

func (s *protoService) Register(grpcServer *grpc.Server) {
	RegisterKeysServer(grpcServer, s)
}

func setupLogging(cfg *Config, logPath string) (Logger, LogInterceptor) {
	return setupLogrus(cfg, logPath)
}

func logFatal(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

// Run the service.
func Run(build Build) {
	appName := flag.String("app", "Keys.pub", "app name")
	logPath := flag.String("log-path", "", "log path")
	version := flag.Bool("version", false, "print version")
	resetKeyring := flag.Bool("reset-keyring", false, "reset keyring")

	flag.Parse()

	if *version {
		fmt.Printf("%s\n", build)
		return
	}

	cfg, err := NewConfig(*appName)
	if err != nil {
		logFatal(errors.Wrapf(err, "failed to load config"))
	}

	if len(flag.Args()) > 0 {
		logFatal(errors.Errorf("Invalid arguments. Did you mean to run `keys`?"))
	}

	if *resetKeyring {
		if IsPortInUse(cfg.Port()) {
			fmt.Printf("Service is running. You should uninstall before resetting.\n")
			os.Exit(1)
		}

		reader := bufio.NewReader(os.Stdin)
		words := keys.RandWords(6)
		fmt.Printf("Are you sure you want to reset the app and remove keys?\n")
		fmt.Printf("If so enter this phrase: %s\n\n", words)
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\r\n")
		fmt.Println("")
		if text != words {
			fmt.Println("Phrase doesn't match.")
			os.Exit(1)
		}
		kr, err := keyring.NewKeyring(cfg.AppName())
		if err != nil {
			logFatal(errors.Wrapf(err, "failed to init keyring"))
		}
		if err := kr.Reset(); err != nil {
			logFatal(errors.Wrapf(err, "failed to reset keyring"))
		}
		fmt.Println("Keyring reset.")
		os.Exit(0)
	}

	lg, lgi := setupLogging(cfg, *logPath)
	SetLogger(lg)
	keys.SetLogger(lg)
	saltpack.SetLogger(lg)
	keyring.SetLogger(lg)
	client.SetLogger(lg)

	if err := checkSupportedOS(); err != nil {
		logFatal(err)
	}
	if runtime.GOOS == "darwin" {
		if err := checkCodesigned(); err != nil {
			logFatal(err)
		}
	}

	logger.Infof("Version: %s", build)
	logger.Infof("Log level: %s", cfg.LogLevel().String())

	panichandler.InstallPanicHandler(func(ctx context.Context, r interface{}) {
		logrus.Errorf("Panic: %v; %s", r, string(debug.Stack()))
	})

	if err := runService(cfg, build, lgi); err != nil {
		logFatal(err)
	}
}

// ServeFn starts the service.
type ServeFn func() error

// CloseFn closes the service.
type CloseFn func()

func runService(cfg *Config, build Build, lgi LogInterceptor) error {
	serveFn, closeFn, serveErr := NewServiceFn(cfg, build, lgi)
	if serveErr != nil {
		return serveErr
	}
	defer closeFn()
	return serveFn()
}

// NewServiceFn ...
func NewServiceFn(cfg *Config, build Build, lgi LogInterceptor) (ServeFn, CloseFn, error) {
	var opts []grpc.ServerOption

	cert, err := generateCA(cfg)
	if err != nil {
		return nil, nil, err
	}
	if cert == nil {
		return nil, nil, errNoCertFound{}
	}
	tlsCert := cert.TLSCertificate()
	creds := credentials.NewServerTLSFromCert(&tlsCert)

	opts = []grpc.ServerOption{
		grpc.Creds(creds),
	}

	auth, err := newAuth(cfg)
	if err != nil {
		return nil, nil, err
	}

	lgi.Replace()

	opts = append(opts, middleware.WithUnaryServerChain(
		ctxtags.UnaryServerInterceptor(ctxtags.WithFieldExtractor(ctxtags.CodeGenRequestFieldExtractor)),
		lgi.Unary(),
		auth.unaryInterceptor,
		panichandler.UnaryServerInterceptor,
	),
		middleware.WithStreamServerChain(
			ctxtags.StreamServerInterceptor(ctxtags.WithFieldExtractor(ctxtags.CodeGenRequestFieldExtractor)),
			lgi.Stream(),
			auth.streamInterceptor,
			panichandler.StreamServerInterceptor,
		),
	)
	grpcServer := grpc.NewServer(opts...)

	service, serviceErr := newProtoService(cfg, build, auth)
	if serviceErr != nil {
		return nil, nil, serviceErr
	}
	service.Register(grpcServer)

	logger.Infof("Listening for connections on port %d", cfg.Port())
	lis, lisErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.Port()))
	if lisErr != nil {
		return nil, nil, errors.Wrapf(lisErr, "failed to tcp listen")
	}

	serveFn := func() error {
		if err := writePID(cfg); err != nil {
			return err
		}
		return grpcServer.Serve(lis)
	}
	closeFn := func() {
		grpcServer.Stop()
		service.Close()
	}
	return serveFn, closeFn, nil
}

// IsPortInUse returns true if port is currently in use.
func IsPortInUse(port int) bool {
	lis, lisErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if lisErr != nil {
		return true
	}
	_ = lis.Close()
	return false
}

func writePID(cfg *Config) error {
	path, err := cfg.AppPath("pid", false)
	if err != nil {
		return err
	}
	pid := os.Getpid()
	return ioutil.WriteFile(path, []byte(strconv.Itoa(pid)), 0600)
}