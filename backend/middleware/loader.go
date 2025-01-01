package middleware

import (
	"net/http"

	"github.com/imaikosuke/deairy/backend/loader"
	"github.com/imaikosuke/deairy/backend/service"
)

func LoaderMiddleware(userSvc service.User, diarySvc service.Diary) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loaders := loader.New(userSvc, diarySvc)
			next.ServeHTTP(w, r.WithContext(loaders.Attach(r.Context())))
		})
	}
}
