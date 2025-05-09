package utils

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(originalPlainPassword string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPlainPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Error Generating hashedPassword")
	}
	return string(hashedPassword), nil
}

func VerifyPasswords(hashedPassword string, userEnteredPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userEnteredPassword))
	if err != nil {
		return errors.New("Incorrect password")
	}
	return nil
}
