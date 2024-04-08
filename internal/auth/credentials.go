package auth

import (
	"encoding/json"
	"errors"
	"go-project/api/models"
	"go-project/internal/database"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SIGNING_KEY"))


// GetCredentials from the request body
func GetCredentials(r *http.Request) (models.Credentials, error) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return models.Credentials{}, err
	}
	return creds, nil
}

// ValidateCredentials checks the credentials and returns user information if valid
func ValidateCredentials(creds models.Credentials) (*models.User, error) {
    user, err := database.GetUserByUsername(creds.Username)
    if err != nil {
        return nil, err
    }

    if !VerifyPasswordHash(creds.Password, user.HashedPassword) {
        return nil, errors.New("invalid credentials")
    }

    return user, nil
}

// CreateSession creates a new session for an authenticated user and returns the session ID
func CreateSession(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	
	claims := &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
		UserID: user.ID,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	sessionToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// In a real application, you might want to store the session information in a store
	// For now, we just return the JWT token string

	return sessionToken, nil
}