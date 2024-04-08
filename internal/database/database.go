package database

import (
	"database/sql"
	"fmt"
	"go-project/api/models"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// DB holds the database connection pool.
var DB *sql.DB

// InitDB initializes the DB variable as a pool of database connections.
func InitDB(psqlDB string) {
	var err error
	DB, err = sql.Open("postgres", psqlDB)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected to the database!")
}

// GetUserByID fetches a user from the database by their ID.
func GetUserByID(id int) (*models.User, error) {
	user := models.User{}

	row := DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}

	row := DB.QueryRow("SELECT id, name, email, hashed_password FROM users WHERE username = $1", username)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user models.UserSignup) error {
	_, err := DB.Exec("INSERT INTO users (name, username, email, hashed_password) VALUES ($1, $2, $3, $4)", user.Name, user.Username, user.Email, user.HashedPassword)
	if err != nil {
		return err
	}

	return nil
}