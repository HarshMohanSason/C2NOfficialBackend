package services

import (
	"c2nofficialsitebackend/utils"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/database"
)

/* --------------- USER SIGN IN ----------------- */
func ProcessUserSignIn(user *models.User) (*models.User, error){
	//First making sure entered user email is in correct format 
	err := utils.ValidateUserEmail(user.Email)
	if err != nil{
		return nil, err
	}
	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}
	//Find the user
	var foundUser *models.User
	foundUser, err = userRepo.SearchUser(user)
	if err != nil{
		return nil, err
	}
	//Verify the passwords now
	if foundUser.Password != nil && user.Password != nil{
	err = utils.VerifyPasswords(*foundUser.Password, *user.Password)
	}
	if err != nil{
		return nil, err
	}	
	return foundUser, nil
}

/* --------------- USER SIGN UP ----------------- */

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