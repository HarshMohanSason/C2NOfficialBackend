package services

import (
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
)

func GoogleAuthService(user *models.User) (*models.User, error) {
	userRepo := database.PostgresUserRepository{DB: database.GetDB()}

	foundUser, err := userRepo.SearchUser(user)
	if foundUser == nil {
		// User not found, create new user
		createErr := userRepo.CreateUser(user)
		if createErr != nil {
			return nil, createErr
		}
		//return the original user since we are signing up with this user
		return user, nil
	}
	if err != nil {
		return nil, err
	}
	// User found, return existing user
	return foundUser, nil
}
