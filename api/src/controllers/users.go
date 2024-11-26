package controllers

import "net/http"

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Adding user..."))
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
