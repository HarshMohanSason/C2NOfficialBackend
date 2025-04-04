package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(originalPlainPassword string) (string, error){

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPlainPassword), bcrypt.DefaultCost)
    if err != nil {
        return "", err 
    }
    return string(hashedPassword), nil
}
