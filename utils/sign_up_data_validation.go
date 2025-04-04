package utils

import (
	"errors"
	"c2nofficialsitebackend/models"
	"github.com/microcosm-cc/bluemonday"
)
//Validating the user info before inserting it into the users table in the db
func ValidateUserInfo(user *models.User) error {

	if user.Name == "" || user.Email == ""{
		return errors.New("Name and user email cannot be empty")
	}
	if user.AuthType == "email"{
		if user.Password == nil{
		return errors.New("Password cannot be empty")
		}
		if len(*user.Password) < 8 || len(*user.Password) > 50{
			return errors.New("Password must be between 8 and 50 characters")
		}
	}
	if len(user.Name) < 2 || len(user.Name) > 50 {
		return errors.New("Name must be between 2 and 50 characters")
	}
	if len(user.Email) < 5 || len(user.Email) > 100 {
		return errors.New("Email must be between 5 and 100 characters")
	}
	//Preventing against XSS attacks. 
	p := bluemonday.StrictPolicy()
	user.Name = p.Sanitize(user.Name)
	user.Email = p.Sanitize(user.Email)

	return nil
}


