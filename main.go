package main

import (
	"day06/internal/db"

	_ "github.com/lib/pq" // че эт за хрень
)

func main() {
	db, err := db.SetupDatabasee()
	if db != nil {
		defer db.Close()
	}
	if err != nil {
		panic(err)
	}

}
