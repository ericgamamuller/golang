package db

import (
	"database/sql"

	_ "github.com/lib/pq"

	"git/challenge-03/config"
)

func Connect() *sql.DB {
	connectionString := config.ReadString("Database.ConnectionString")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return db
}
