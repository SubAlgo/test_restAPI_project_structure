package configDB

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Postgres() (db *sql.DB) {
	connStr := "postgres://postgres:test123456@localhost:8100/test_api?sslmode=disable"
	dbURL := os.Getenv("check_cdd_db_url")

	if dbURL == "" {
		//dbURL = "postgres://localhost:6900/cdd?sslmode=disable"
		dbURL = connStr
	}
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("can not connect to database;", err)
	}
	return
}
