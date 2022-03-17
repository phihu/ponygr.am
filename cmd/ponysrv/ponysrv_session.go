package main

import (
	"context"
	"net/http"

	"github.com/phihu/ponygr.am/pkg/log"
	"github.com/phihu/ponygr.am/pkg/session"
	"github.com/pkg/errors"
)

func SessionMiddleware(cfg Config, sessionSvc session.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			sess, err := sessionSvc.Store().Get(r, cfg.SessionName())
			if err != nil {
				if errors.Cause(err) != context.Canceled {
					// nolint: errcheck
					log.ErrLog(ctx).Log("msg", "error loading session",
						"err", err,
					)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			ctx = session.ContextWithSession(ctx, sess)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
