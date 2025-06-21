package middlewares

import (
	"context"
	"day06/internal/rate_limiter"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func RateLimiterMiddleware(limiter *rate_limiter.RateLimiter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.Background()
			key := r.RemoteAddr

			// For testing change to 10 requests per minute
			res, err := limiter.Limit(ctx, key, 100, time.Second)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if res.Allowed == 0 {
				w.Header().Set("Retry-After", res.RetryAfter.String())
				http.Error(w, "429 Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
