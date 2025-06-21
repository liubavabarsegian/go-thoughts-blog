package app

import (
	"database/sql"
	"day06/internal/app/middlewares"
	"day06/internal/app/repository"
	"day06/internal/app/router"
	"day06/internal/app/service"
	"day06/internal/db"
	"day06/internal/rate_limiter"
	"net/http"

	"github.com/gorilla/mux"
)

// set up logger
type App struct {
	RateLimiter *rate_limiter.RateLimiter
	DB          *sql.DB
}

func New() *App {
	db, err := db.SetupDatabase()
	if err != nil {
		return &App{}
	}

	rate_limiter := rate_limiter.NewRateLimiter()
	return &App{DB: db, RateLimiter: rate_limiter}
}

func (app *App) Go() (err error) {
	if app.DB == nil || app.RateLimiter == nil {
		return
	}
	defer app.DB.Close()

	// Создание репозиториев
	repo := repository.NewRepository(app.DB)
	// Создание сервисного слоя
	services := service.NewService(repo)
	// Настройка роутера
	mux := mux.NewRouter()
	router.SetupHandlers(mux, services)

	mux.Use(middlewares.WithLogger)
	mux.Use(middlewares.PanicMiddleware)
	mux.Use(middlewares.RateLimiterMiddleware(app.RateLimiter))

	http.ListenAndServe(":8888", mux)
	return
}
