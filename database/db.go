package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func init() {
	DB = setup()
}

func setup() *sql.DB {
	db, err := sql.Open("sqlite3", "../database/db.sqlite")
	if err != nil {
		log.Panic(err.Error())
	}
	log.Print("DB connected")
	return db
}
