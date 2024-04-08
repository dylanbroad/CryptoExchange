package main

import (
	"fmt"
	"go-project/api/handlers"
	"go-project/internal/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file:", err)
    }

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"))

	// Initialize the database connection.
	database.InitDB(psqlconn)
	defer database.DB.Close()

	// Create a new router
	r := mux.NewRouter()

	// Setup route handlers
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/auth/signin", handlers.SignIn).Methods("POST")
	// ... more routes ...

	// Apply middleware
	// r.Use(middleware.YourMiddleware)

	certPath := "server.crt"
	keyPath := "server.key"

	// Start the server
	log.Println("Starting server on :8080")
	err := http.ListenAndServeTLS(":8080", certPath, keyPath, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
