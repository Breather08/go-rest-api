package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var err error
var db *sql.DB
var dbFileName = "database.db"

func InitDB() {
	os.Remove(dbFileName)

	file, err := os.Create(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	fmt.Printf("%v created successfully\n", dbFileName)

	openDB()
	createTables()
}

func createTables() {
	CreatePostTable()
}

func openDB() {
	db, err = sql.Open("sqlite3", "./"+dbFileName)
	if err != nil {
		log.Fatal(err)
	}
}

func CloseDB() {
	db.Close()
}
