package app

import (
	"database/sql"
	"day06/internal/app/middlewares"
	"day06/internal/app/repository"
	"day06/internal/app/router"
	"day06/internal/app/service"
	"day06/internal/db"
	"net/http"

	"github.com/gorilla/mux"
	// "honnef.co/go/tools/config"
)

// set up envs
// set up logger
type App struct {
	// Config *config.Config
	Server *http.Server
	DB     *sql.DB
}

func New() *App {
	return &App{}
}

func (app *App) Go() (err error) {
	db, err := db.SetupDatabase()
	if err != nil || db == nil {
		return
	}
	defer db.Close()

	// Создание репозиториев
	repo := repository.NewRepository(db)
	// Создание сервисного слоя
	services := service.NewService(repo)

	mux := mux.NewRouter()
	router.SetupHandlers(mux, services)
	mux.Use(middlewares.WithLogger)

	http.ListenAndServe(":8888", mux)
	return
}
