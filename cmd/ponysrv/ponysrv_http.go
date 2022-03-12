package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/phihu/ponygr.am/pkg/session"
	"github.com/phihu/ponygr.am/pkg/user"
)

type (
	Config interface {
		SessionName() string
	}
)

func DefaultMiddlewareChain(
	cfg Config,
	sessionSvc session.Service) chi.Middlewares {
	mws := []func(http.Handler) http.Handler{
		SessionMiddleware(cfg, sessionSvc),
	}
	return chi.Chain(mws...)
}

func Mount(r chi.Router,
	cfg Config,
	sessionSvc session.Service,
	userSvc user.Service,
) error {
	chain := DefaultMiddlewareChain(cfg, sessionSvc)
	r = r.With(chain...)

	r.Route("/user", func(r chi.Router) {
		// auth required
		r.Group(func(r chi.Router) {
			r = r.With(UserRequireMiddleware(userSvc))
			r.Get("/", UserGetHandler(userSvc))
		})
	})

	return nil
}
