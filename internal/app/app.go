package app

import (
	"database/sql"
	"day06/internal/app/middlewares"
	"day06/internal/app/router"
	"day06/internal/db"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
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
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &App{}
}

func (app *App) Go() (err error) {
	fmt.Println("started")
	db, err := db.SetupDatabase()
	if err != nil || db == nil {
		return
	}
	defer db.Close()

	mux := app.SetUpRouter()
	router.SetupHandlers(mux)
	mux.Use(middlewares.WithLogger)

	// app.Logger.Info().Msgf("Server is listening: %s", "8888")
	http.ListenAndServe(":8888", mux)
	return
}

func (app *App) SetUpRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}
