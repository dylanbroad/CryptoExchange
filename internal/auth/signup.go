package auth

import (
	"encoding/json"
	"net/http"
	"go-project/api/models"
)


func GetSignUpCreds (r *http.Request) (models.UserSignup, error) {
	var creds models.UserSignup
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return models.UserSignup{}, err
	}
	return creds, nil
}