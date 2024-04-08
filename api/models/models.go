// model.go

package models

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	ID    int     `json:"id"`
    Name  string  `json:"name"`
	Username  string  `json:"username"`
    Email string  `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

type UserSignup struct {
	Name string `json:"name"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	HashedPassword string `json:"hashed_password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type Wallet struct {
	WalletID int `json:"wallet_id"`
	UserID int `json:"user_id"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	TransactionID int `json:"transaction_id"`
	WalletID int `json:"wallet_id"`
	Amount float64 `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Type string `json:"type"`
	Status string `json:"status"`
}
