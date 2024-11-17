package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plaintext string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plaintext), 10)
	return string(hashed), err
}

func ValidatePassword(password, hashedPassword string) (bool, error) {
	hashedPass, plaintext := []byte(hashedPassword), []byte(password)
	if err := bcrypt.CompareHashAndPassword(hashedPass, plaintext); err != nil {
		return false, err
	} 
	return true, nil
}