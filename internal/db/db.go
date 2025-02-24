package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "username"
	password = "password"
	dbname   = "dbname"
)

func SetupDatabase() (*sql.DB, error) {
	db, err := connectDatabase()
	if err != nil {
		return nil, fmt.Errorf("error while connection database", err)
	}

	err = migrateDatabase(db)
	if err != nil {
		return nil, fmt.Errorf("error while migrating database", err)
	}
	return db, nil
}

func connectDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(password)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to DB", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error while pinging DB", err)
	}

	return db, nil
}

func migrateDatabase(db *sql.DB) error {
	wd, err := os.Getwd()
	fmt.Println(wd)
	sqlFile, err := os.ReadFile("/app/internal/db/schema.sql")
	if err != nil {
		return fmt.Errorf("error while migrating DB schema", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("error while migrating DB schema", err)
	}
	return nil
}
