package server

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/keys-pub/keys"
	"github.com/pkg/errors"
)

// AuthResult is the authorized result.
type AuthResult struct {
	kid    keys.ID
	spk    keys.SignPublicKey
	method string
	url    *url.URL
	nonce  string
	ts     time.Time
}

// CheckAuthorization returns error if authorization fails.
// The auth is of the form:
// <PKID>:<SIG>
// The <SIG> is the detached signature of <METHOD>,<URL>.
func CheckAuthorization(ctx context.Context, method string, urs string, auth string, mc MemCache, now time.Time) (*AuthResult, error) {
	fields := strings.Split(auth, ":")
	if len(fields) != 2 {
		return nil, errors.Errorf("too many fields")
	}
	skid := fields[0]
	sig := fields[1]

	kid, err := keys.ParseID(skid)
	if err != nil {
		return nil, err
	}

	spk, spkerr := keys.DecodeSignPublicKey(kid.String())
	if spkerr != nil {
		return nil, errors.Errorf("not a valid sign public key")
	}

	sigBytes, sigerr := keys.Decode(sig, keys.Base62)
	if sigerr != nil {
		return nil, sigerr
	}

	url, err := url.Parse(urs)
	if err != nil {
		return nil, err
	}

	msg := method + "," + url.String()
	logger.Infof(ctx, "Checking auth for %s %s", msg, auth)
	if err := keys.VerifyDetached(sigBytes, []byte(msg), spk); err != nil {
		return nil, err
	}

	nonce := url.Query().Get("nonce")
	if nonce == "" {
		return nil, errors.Errorf("nonce is missing")
	}
	nb, err := keys.Decode(nonce, keys.Base62)
	if err != nil {
		return nil, err
	}
	if len(nb) != 32 {
		return nil, errors.Errorf("nonce is invalid length")
	}
	// Namespace the nonce key
	nonce = fmt.Sprintf("auth-%s", nonce)

	val, err := mc.Get(ctx, nonce)
	if err != nil {
		return nil, err
	}
	if val != "" {
		return nil, errors.Errorf("nonce collision")
	}
	if err := mc.Set(ctx, nonce, "1"); err != nil {
		return nil, err
	}
	if err := mc.Expire(ctx, nonce, time.Hour); err != nil {
		return nil, err
	}

	// Check timestamp
	ts := url.Query().Get("ts")
	if ts == "" {
		return nil, errors.Errorf("timestamp (ts) is missing")
	}
	i, err := strconv.Atoi(ts)
	if err != nil {
		return nil, err
	}
	tm := keys.TimeFromMillis(keys.TimeMs(i))
	td := now.Sub(tm)
	if td < 0 {
		td = td * -1
	}
	if td > 30*time.Minute {
		return nil, errors.Errorf("timestamp is invalid, diff %s", td)
	}

	return &AuthResult{
		kid:    kid,
		spk:    spk,
		method: method,
		url:    url,
		nonce:  nonce,
		ts:     tm,
	}, nil

}