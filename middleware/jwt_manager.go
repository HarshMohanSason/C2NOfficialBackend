
package middleware
/*
import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"net/http"
)

func VerifyJWT(handlerFunc http.Handler) http.Handler{
	return http.HandlerFunc(func(response http.ResponseWriter, receivedRequest http.Request){

		cookie, err := receivedRequest.Cookie("token")
		if err != nil{
			http.Error(response, "Unauthorized request, please try again", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error){
			return os.Getenv("JWT_SECRET"), nil
		})
		if err != nil{
			http.Error(response, err.Error(), http.StatusUnauthorized)
			return
		}
		handlerFunc.ServeHTTP(receivedRequest, response)
	})
}

func GenerateJWT(username String) (string, error){

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
    	jwt.MapClaims{ 
    	"username": username, 
    	"exp": time.Now().Add(time.Hour * 24).Unix(), 
    })
    tokenString, err := token.SignedString(os.Getenv("JWT_SECRET")) //Sign the token with the secret key
    if err != nil {
    	return "", err
    }
    return tokenString, nil
}
*/

