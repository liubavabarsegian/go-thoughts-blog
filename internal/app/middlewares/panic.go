package middlewares

import (
	"fmt"
	"net/http"
)

// func accessLogMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("accessLogMiddleware", r.URL.Path)
// 		start := time.Now()
// 		next.ServeHTTP(w, r)
// 		fmt.Printf("[%s] %s, %s %s\n",
// 			r.Method, r.RemoteAddr, r.URL.Path, time.Since(start))
// 	})
// }

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
