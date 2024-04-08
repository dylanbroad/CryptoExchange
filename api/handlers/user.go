package handlers

import (
	"encoding/json"
	"go-project/internal/auth"
	"go-project/internal/database"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser handles the POST request to create a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	creds, err := auth.GetSignUpCreds(r)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }
    creds.HashedPassword, err = auth.HashPassword(creds.Password)
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = database.CreateUser(creds)
    if err != nil {
        http.Error(w, "Failed to Create User", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := database.GetUserByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
