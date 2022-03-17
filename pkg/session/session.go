package session

import (
	"context"

	"github.com/gorilla/sessions"
)

type (
	contextKey int
)

const (
	contextKeySession contextKey = iota
)

func FromContext(ctx context.Context) *sessions.Session {
	if ctx.Value(contextKeySession) == nil {
		return nil
	}
	if sess, ok := ctx.Value(contextKeySession).(*sessions.Session); ok {
		return sess
	}
	return nil
}

func ContextWithSession(ctx context.Context, sess *sessions.Session) context.Context {
	return context.WithValue(ctx, contextKeySession, sess)
}
