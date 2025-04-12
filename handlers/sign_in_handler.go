package handlers

import(
	"net/http"
	"io"
	"encoding/json"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/services"
	"c2nofficialsitebackend/middleware"
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
	var returnedUser *models.User
	returnedUser, err = services.ProcessUserSignIn(&user); 

	if err != nil{
		http.Error(response,err.Error(),http.StatusConflict)
		return
	}
	//Do not need the error since a jwt not being generated should be ignored
	tokenJWT, _ := middleware.GenerateJWT(returnedUser.Name) 
		
	//Set cookie for JWT 
	http.SetCookie(response, &http.Cookie{
	  Name:     "auth-token",
	  Value:    tokenJWT,
	  HttpOnly: true,
	  Secure:   false,  
	  Path:     "/",
	  SameSite: http.SameSiteLaxMode,
	})

	//Set cookie for email
	http.SetCookie(response, &http.Cookie{
	  Name:     "email",
	  Value:    returnedUser.Email,
	  HttpOnly: true,
	  Secure:   false,  
	  Path:     "/",
	  SameSite: http.SameSiteLaxMode,
	})

	//Set cookie for auth type
	http.SetCookie(response, &http.Cookie{
	  Name:     "auth-type",
	  Value:    returnedUser.AuthType,
	  HttpOnly: true,
	  Secure:   false,  
	  Path:     "/",
	  SameSite: http.SameSiteLaxMode,
	})
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("User signed in successfully"))
}