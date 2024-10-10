package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	fmt.Println(user)
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