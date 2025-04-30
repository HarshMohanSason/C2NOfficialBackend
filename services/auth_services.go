package services

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/utils"
	"errors"
)

func ProcessUserSignIn(user *models.User) (*models.User, error) {

	//making sure entered user email is in correct format
	err := utils.ValidateUserEmail(user.Email)
	if err != nil {
		return nil, err
	}
	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}
	//Find the user with that email
	var foundUser *models.User
	foundUser, _ = userRepo.SearchUser(user)

	//Proceed to signup if none is found with Google login
	if foundUser == nil && user.AuthType == "google" {
		err = ProcessUserSignUp(user)
		if err != nil {
			return nil, err
		}
		foundUser = user //Found user is the new user
	}

	if foundUser == nil {
		return nil, errors.New("no user not found")
	}

	//Verify the passwords if the user signed in via email
	if user.AuthType == "email" && user.Password != nil {
		if err = utils.VerifyPasswords(*foundUser.Password, *user.Password); err != nil {
			return nil, err
		}
	}
	//Once the user is found, set the current auth roles for that user.
	setRoleError := database.SetUserRole(database.GetDB(), foundUser)
	if setRoleError != nil {
		return nil, setRoleError
	}
	return foundUser, nil
}

func ProcessUserSignUp(user *models.User) error {
	if err := validateAndSanitizeUser(user); err != nil {
		config.LogError(err)
		return err
	}
	if err := saveUserToRepository(user); err != nil {
		config.LogError(err)
		return err
	}
	//Once the user is signed up, set the current auth roles for that user.
	setRoleError := database.SetUserRole(database.GetDB(), user)
	if setRoleError != nil {
		return nil
	}
	return nil
}

// startSanitizingData validates and hashes the user's password
func validateAndSanitizeUser(user *models.User) error {
	if err := utils.ValidateUserInfo(user); err != nil {
		return err
	}
	if user.Password != nil {
		hashedPassword, err := utils.GenerateHashedPassword(*user.Password)
		if err != nil {
			return err
		}
		user.Password = &hashedPassword
	}
	return nil
}

// StartSavingUserToDB saves user to database
func saveUserToRepository(user *models.User) error {
	db := database.GetDB()                               //Get the db instance
	userRepo := &database.PostgresUserRepository{DB: db} //Instance to create userRepo
	err := createNewUser(userRepo, user)                 //Create the user
	if err != nil {
		return err
	}
	return nil
}

// CreateNewUser adheres to the UserRepository interface
func createNewUser(userRepo database.UserRepository, user *models.User) error {
	return userRepo.CreateUser(user) // Return error directly
}
