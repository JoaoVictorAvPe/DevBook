package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepositorie := repositories.NewUserRepositorie(db)
	userId, err := userRepositorie.Add(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("User inserted, id: %d", userId)))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users..."))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting user..."))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Uptating user..."))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
