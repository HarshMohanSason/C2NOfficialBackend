package utils

import (
	"c2nofficialsitebackend/models"
	"errors"
	"github.com/microcosm-cc/bluemonday"
	"regexp"
)

// ValidateUserInfo before inserting it into the users table in the db
func ValidateUserInfo(user *models.User) error {

	err := ValidateUserEmail(user.Email)
	if err != nil {
		return err
	}
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	if len(user.Name) < 2 || len(user.Name) > 50 {
		return errors.New("name must be between 2 and 50 characters")
	}
	if user.AuthType == "email" {
		if user.Password == nil {
			return errors.New("password cannot be empty")
		}
		if len(*user.Password) < 8 || len(*user.Password) > 50 {
			return errors.New("password must be between 8 and 50 characters")
		}
	}

	//Preventing against XSS attacks.
	p := bluemonday.StrictPolicy()
	user.Name = p.Sanitize(user.Name)
	user.Email = p.Sanitize(user.Email)

	return nil
}

func ValidateUserEmail(email string) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	didMatch, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return err
	}
	if !didMatch {
		return errors.New("invalid email format")
	}
	return nil
}
