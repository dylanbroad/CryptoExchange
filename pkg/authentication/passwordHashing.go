package passwordHashing

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword (password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func VerifyPasswordHash (password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main () {
    password := "myP@ssw0rd!"
    hashedPassword, err := HashPassword(password)
    if err != nil {
        panic(err)
    }
    fmt.Println("Hashed Password:", hashedPassword)

    passwordAttempt := "myP@ssw0rd!"
	if VerifyPasswordHash(passwordAttempt, hashedPassword) {
		fmt.Println("Password is correct!")
	} else {
		fmt.Println("Password is incorrect!")
	}	
}
