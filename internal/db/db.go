package db

import (
	"database/sql"
	"fmt"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "liuba"
	password = "password"
	dbname   = "blog_database"
)

func SetupDatabasee() (*sql.DB, error) {
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
	sqlFile, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		return fmt.Errorf("error while migrating DB schema", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return fmt.Errorf("error while migrating DB schema", err)
	}
	return nil
}
