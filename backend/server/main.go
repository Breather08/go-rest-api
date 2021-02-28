package main

import (
	"fmt"
	"forum/backend/server/database"
	"forum/backend/server/router"
	"log"
	"net/http"
)

func init() {
	database.InitDB()
}

func main() {
	defer func() {
		database.CloseDB()
		fmt.Println("DB closed")
	}()

	router := router.Router()

	fmt.Println("Starting server on port 7880...")

	log.Fatal(http.ListenAndServe(":7880", router))
}
