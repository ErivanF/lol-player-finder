package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func OpenConnection() (*sql.DB, error) {
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	conn, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" password="+password+" dbname="+dbname+" sslmode=disable")
	if err != nil {
		panic(err)
	}
	Conn = conn
	err = Conn.Ping()
	if err != nil {
		panic(err)
	}
	return Conn, err
}
