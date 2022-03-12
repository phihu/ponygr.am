package main

import (
	"net/http"

	"github.com/phihu/ponygr.am/pkg/user"
)

func UserRequireMiddleware(userSvc user.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})
	}
}

func UserGetHandler(userSvc user.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}
