package middlewares

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// WithLogger — middleware для логирования запросов
func WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем запрос
		logger := log.Ctx(r.Context())
		// logger := log.With().Str("request_id", reqID).Caller().Logger()
		ctx := logger.WithContext(r.Context())
		logger.Printf("[%s] %s %s", r.Method, r.URL.Path, time.Since(start))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
