package app

import (
	"day06/internal/app/router"
	"day06/internal/db"
	"log"
	"net/http"
)

func Go() error {
	// set up envs
	//set up logger
	db, err := db.SetupDatabasee()
	if db != nil {
		defer db.Close()
	}
	if err != nil || db == nil {
		return err
	}
	// set up router
	mux := router.SetUpRouter()
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, mux)

	return nil
}
