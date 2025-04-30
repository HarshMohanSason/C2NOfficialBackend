package authHandlers

import (
	"c2nofficialsitebackend/middleware"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
	"c2nofficialsitebackend/utils"
	"encoding/json"
	"io"
	"net/http"
)

func EmailSignInHandler(response http.ResponseWriter, receivedRequest *http.Request) {

	if receivedRequest.Method != http.MethodPost {
		http.Error(response, "Invalid request method", http.StatusMethodNotAllowed)
	}

	body, err := io.ReadAll(receivedRequest.Body)

	if err != nil {
		http.Error(response, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer receivedRequest.Body.Close()

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		http.Error(response, "Invalid Format, please try again", http.StatusBadRequest)
		return
	}
	var returnedUser *models.User
	returnedUser, err = services.SignInWithEmail(&user)

	if err != nil {
		http.Error(response, err.Error(), http.StatusConflict)
		return
	}
	tokenJWT, _ := middleware.GenerateJWT(returnedUser.Name)

	utils.SetAuthCookies(response,
		&utils.Cookie{Name: "auth-token", Value: tokenJWT, Path: "/"},
		&utils.Cookie{Name: "email", Value: user.Email, Path: "/"},
		&utils.Cookie{Name: "auth-type", Value: user.AuthType, Path: "/"},
	)
}
