package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Conn *sql.DB

func OpenDatabase() {
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	credentials := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname = %s sslmode=disable", host, user, password, name)
	var err error
	Conn, err = sql.Open("postgres", credentials)
	if err != nil {
		panic(err)
	}
}

func MigrateUp() {
	driver, err := postgres.WithInstance(Conn, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Up()
}

func GetDb() *sql.DB {
	return Conn
}
