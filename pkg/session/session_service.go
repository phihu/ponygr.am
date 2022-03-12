package session

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/gorilla/sessions"
)

const SignatureKeySize = 64

type (
	Service interface {
		Opts() *sessions.Options
		Store() sessions.Store
		New(ctx context.Context, sess *sessions.Session) (*sessions.Session, error)
		Regenerate(ctx context.Context, sess *sessions.Session) error
	}
	Config interface {
		SessionSignatureKeys() [][]byte
		HTTPSecure() bool
		SessionTimeout() time.Duration
	}
	service struct {
		cfg   Config
		store sessions.Store
	}
)

func NewService(cfg Config) (Service, error) {
	svc := &service{cfg: cfg}
	storeKeys := make([][]byte, 0, 4)
	keys := cfg.SessionSignatureKeys()
	if len(keys) == 0 {
		return nil, errors.New("missing signature key")
	}
	if len(keys[0]) != SignatureKeySize {
		return nil, fmt.Errorf("expect signature key to be %d, have %d", SignatureKeySize, len(keys[0]))
	}
	storeKeys = append(storeKeys, keys[0], nil)
	if len(keys) == 2 {
		if len(keys[1]) != SignatureKeySize {
			return nil, fmt.Errorf("expect signature key to be %d, have %d", SignatureKeySize, len(keys[1]))
		}
		storeKeys = append(storeKeys, keys[1], nil)
	}
	svc.store = sessions.NewCookieStore(storeKeys...)
	return svc, nil
}

func (svc *service) Opts() *sessions.Options {
	return &sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   int(svc.cfg.SessionTimeout()),
		Secure:   svc.cfg.HTTPSecure(),
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}
}

func (svc *service) Store() sessions.Store {
	return svc.store
}

func (svc *service) New(ctx context.Context, sess *sessions.Session) (*sessions.Session, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return sess, errors.Wrap(err, "error generating new session ID")
	}

	sess.Options = svc.Opts()
	sess.ID = id.String()
	sess.IsNew = true
	return sess, nil
}

func (svc *service) Regenerate(ctx context.Context, sess *sessions.Session) error {
	id, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "error generating session ID")
	}
	sess.ID = id.String()
	sess.IsNew = true
	return nil
}
