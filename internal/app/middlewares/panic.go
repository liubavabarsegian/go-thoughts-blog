package middlewares

import (
	"fmt"
	"net/http"
)

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panicMiddleware: recovered", r.URL.Path, err)
				http.Error(w, "panicMiddleware: Internal server error", 500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
