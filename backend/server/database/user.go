package database

import (
	"encoding/json"
	"fmt"
	"forum/backend/server/models"
	"log"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const cookieName = "session_token"

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

func authenticateUser(email string, password string) int {
	var id int
	var hashedPassword string

	row := db.QueryRow(getUserSQL, email)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		fmt.Println("Error occured while scanning through db", err)
		return 0
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("Invalid credentials")
		return 0
	}
	return id
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// cookie, err := r.Cookie(cookieName)
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// sessionToken = cookie.Value

	
}

func setCookie(w http.ResponseWriter) {
	sessionToken := uuid.NewV4().String()
	
	cookie := &http.Cookie{
		Name: cookieName,
		Value: sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	}

	http.SetCookie(w, cookie)
}

func GetUser(id int) models.User {
	var user models.User
	return user
}
