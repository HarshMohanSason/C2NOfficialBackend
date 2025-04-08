
package handlers

import(
	"net/http"
	"io"
	"encoding/json"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
)

func ReceiveSignInFormUserInfo(response http.ResponseWriter, receivedRequest *http.Request){

	if receivedRequest.Method != http.MethodPost{
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

	if err != nil{
		http.Error(response, "Invalid Format, please try again", http.StatusBadRequest)
		return
	}
	err = services.ProcessUserSignIn(&user); 

	if err != nil{
		http.Error(response,err.Error(),http.StatusConflict)
		return
	}
}