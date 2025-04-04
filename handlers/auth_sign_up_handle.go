package handlers

import (
	"net/http"
	"io"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
	"encoding/json"
)

func ReceiveSignUpFormUserInfo(response http.ResponseWriter, receivedRequest *http.Request){

	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Pre flight request 
	if receivedRequest.Method == http.MethodOptions{
		response.WriteHeader(http.StatusOK)
		return 
	}

	if receivedRequest.Method != http.MethodPost{
		http.Error(response, "Invalid request", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(receivedRequest.Body)

	if err != nil {
		http.Error(response, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer receivedRequest.Body.Close() //Close the body once the function finishes

	var user models.User

	//Create a User obejct from the received JSON 
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(response, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//Start processing the user 
	err = services.ProcessUserSignUp(&user)
	if err != nil{
		http.Error(response, err.Error(), http.StatusConflict)
		return
	}
}
