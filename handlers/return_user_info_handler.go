package handlers

import (
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/models"
	"encoding/json"
	"net/http"
)

func ReturnUserInfo(response http.ResponseWriter, receivedRequest *http.Request) {

	if receivedRequest.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userRepo := &database.PostgresUserRepository{DB: database.GetDB()}

	authTokenCookie, err := receivedRequest.Cookie("auth-token")
	if err != nil {
		return // No cookie found, just return
	}

	emailCookie, err := receivedRequest.Cookie("email")
	if err != nil {
		return
	}

	authTypeCookie, err := receivedRequest.Cookie("auth-type")
	if err != nil {
		return
	}

	user := models.User{
		Name:     authTokenCookie.Value,
		Email:    emailCookie.Value,
		AuthType: authTypeCookie.Value,
	}

	foundUser, _ := userRepo.SearchUser(&user)
	if foundUser == nil {
		return
	}

	response.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(response).Encode(foundUser); err != nil {
		http.Error(response, "Failed to encode user", http.StatusInternalServerError)
	}
}
