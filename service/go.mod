module github.com/keys-pub/keysd/service

go 1.12

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/gogo/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/keys-pub/keys v0.0.0-20200128210451-9ae521aecfd0
	github.com/keys-pub/keysd/db v0.0.0-20200110200550-f472e79d2ba3
	github.com/keys-pub/keysd/http/api v0.0.0-20200128211114-8182e2f364bf
	github.com/keys-pub/keysd/http/client v0.0.0-20200128211114-8182e2f364bf
	github.com/keys-pub/keysd/http/server v0.0.0-20200128211114-8182e2f364bf
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/logrusorgru/aurora v0.0.0-20191116043053-66b7ad493a23
	github.com/mercari/go-grpc-interceptor v0.0.0-20180110035004-b8ad3827e82a
	github.com/mitchellh/go-ps v0.0.0-20190716172923-621e5597135b
	github.com/onsi/ginkgo v1.9.0 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.4.0
	github.com/urfave/cli v1.22.2
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	google.golang.org/genproto v0.0.0-20191216205247-b31c10ee225f // indirect
	google.golang.org/grpc v1.26.0
)

// replace github.com/keys-pub/keys => ../../keys

// replace github.com/keys-pub/keysd/db => ../db

// replace github.com/keys-pub/keysd/http/api => ../http/api

// replace github.com/keys-pub/keysd/http/client => ../http/client

// replace github.com/keys-pub/keysd/http/server => ../http/server
