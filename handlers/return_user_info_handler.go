package handlers

import (
	"net/http"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/database"
	"encoding/json"
)

func ReturnUserInfo(response http.ResponseWriter, receivedRequest *http.Request){

	if receivedRequest.Method != http.MethodGet{
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
	}

	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}
	
	//Not handling any errors here since it is being already verified with the jwt verify middleware
	authTokenCookie, _ := receivedRequest.Cookie("auth-token")
	emailCookie, _ := receivedRequest.Cookie("email")
	authTypeCookie, _ := receivedRequest.Cookie("auth-type")
	
	//Creating basic info for the user to search for the user
	user := models.User{
		Name: authTokenCookie.Value, 
		Email: emailCookie.Value, 
		AuthType: authTypeCookie.Value,
	}

	foundUser, err := userRepo.SearchUser(&user)
	if err != nil{
		http.Error(response, "Could not find the user, please login again to continue", http.StatusUnauthorized)
		return 
	}
	//Sending the response back now
	response.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(response).Encode(foundUser); err != nil {
	http.Error(response, "Failed to encode user", http.StatusInternalServerError)
	}
}
