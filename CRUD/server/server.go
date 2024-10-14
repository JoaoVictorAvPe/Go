package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID		uint32	`json:"id"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`	

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read the request body"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Failed to precess the request body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect on database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		w.Write([]byte("Failed to create statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Failed to insert user"))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to obtain id inserted"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User %s has been inserted sucessfuly. ID: %d", user.Name, idInserted)))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	lines, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.Write([]byte("Failed to get users from database"))
		return
	}
	defer lines.Close()

	var users []user
	for lines.Next() {
		var userTemp user

		if err := lines.Scan(&userTemp.ID, &userTemp.Name, &userTemp.Email); err != nil {
			w.Write([]byte("Failed to scan users"))
		}

		users = append(users, userTemp)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Failed to convert data to json"))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	Id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to convert param to integer"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
	}

	line, err := db.Query("SELECT * FROM usuarios WHERE id = ?", Id)
	if err != nil {
		w.Write([]byte("Failed to get values from database"))
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Failed to scan values from database"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Failed to convert user to json"))
		return
	}

}