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

func EmailSignUpHandler(response http.ResponseWriter, receivedRequest *http.Request) {

	if receivedRequest.Method != http.MethodPost {
		http.Error(response, "Invalid request", http.StatusMethodNotAllowed)
		return
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

	err = services.SignUpWithEmail(&user)
	if err != nil {
		http.Error(response, err.Error(), http.StatusConflict)
		return
	}

	//Do not need the error since a jwt not being generated should be ignored
	tokenJWT, _ := middleware.GenerateJWT(user.Name)

	utils.SetAuthCookies(response,
		&utils.Cookie{Name: "auth-token", Value: tokenJWT, Path: "/"},
		&utils.Cookie{Name: "email", Value: user.Email, Path: "/"},
		&utils.Cookie{Name: "auth-type", Value: user.AuthType, Path: "/"},
	)
}
