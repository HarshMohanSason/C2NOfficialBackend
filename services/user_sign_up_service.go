package services

import (
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/utils"
)

func ProcessUserSignUp(user *models.User) error {
	if err := validateAndSanitizeUser(user); err != nil {
		utils.LogError(err)
		return err
	}
	if err := saveUserToRepository(user); err != nil {
		utils.LogError(err)
		return err
	}
	return nil
}

// startSanitizingData validates and hashes the user's password
func validateAndSanitizeUser(user *models.User) error {
	if err := utils.ValidateUserInfo(user); err != nil {
		return err
	}
	if user.Password != nil{
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

	db := database.GetDB() //Get the db instance
	userRepo := &database.PostgresUserRepository{DB: db} //Instance to create userRepo 
	err := createNewUser(userRepo, user) //Create the user 
	if err != nil{
		return err
	}
	return nil
}

// CreateNewUser adheres to the UserRepository interface
func createNewUser(userRepo database.UserRepository, user *models.User) error {
	return userRepo.CreateUser(user) // Return error directly
}