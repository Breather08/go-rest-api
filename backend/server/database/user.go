package database

import (
	"encoding/json"
	"fmt"
	"forum/backend/server/models"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUsersTable() {
	userTable, err := db.Prepare(userTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = userTable.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Users table created in DB")
}

func insertUserIntoDB(user models.User) {
	insertPost, err := db.Prepare(insertUserSQL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = insertPost.Exec(
		user.Username,
		user.Email,
		user.HashedPassword,
		user.CreatedAt,
		user.Active,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New user inserted into DB")
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), 12)
    if err != nil {
		log.Println(err)
    }
	
	user.CreatedAt = time.Now()
	user.HashedPassword = string(hashedPassword)
	user.Active = true
	// Check for duplicate parameters in DB
	insertUserIntoDB(user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New user sent to clientside")
}

func AuthenticateUser(email string, password string) int {
	return 0
}

func GetUser(id int) models.User {
	var user models.User
	return user
}
