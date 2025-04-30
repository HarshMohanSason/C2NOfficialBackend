package services

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/utils"
)

func SignInWithEmail(user *models.User) (*models.User, error) {
	//make sure entered user email is in correct format
	err := utils.ValidateUserEmail(user.Email)
	if err != nil {
		return nil, err
	}
	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}
	//Search for that user with that mail
	foundUser, err := userRepo.SearchUser(user)
	if err != nil {
		config.LogError(err)
		return nil, err
	}
	//If a user is found with that email, now verifying the passwords
	if user.Password != nil {
		err = utils.VerifyPasswords(*foundUser.Password, *user.Password)
		if err != nil {
			config.LogError(err)
			return nil, err
		}
	}
	//Set the role for the user(If the user is admin)
	setRoleError := database.SetUserRole(database.GetDB(), foundUser)
	if setRoleError != nil {
		config.LogError(setRoleError)
		return nil, setRoleError
	}
	return foundUser, nil
}

func SignUpWithEmail(user *models.User) error {
	//validate the name and email
	err := utils.ValidateUserInfo(user)
	if err != nil {
		config.LogError(err)
		return err
	}
	//Generate a hash for the user entered password text
	if user.Password != nil {
		hashedPassword, err := utils.GenerateHashedPassword(*user.Password)
		if err != nil {
			config.LogError(err)
			return err
		}
		user.Password = &hashedPassword
	}
	//start saving the user info to db
	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}
	err = userRepo.CreateUser(user)
	if err != nil {
		config.LogError(err)
		return err
	}
	return nil
}
