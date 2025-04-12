
package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"net/http"
	"os"
	"c2nofficialsitebackend/utils"
)

/* Not handling errors explicitly in the entire JWT file. Only
logging the errors
*/ 

func VerifyJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth-token")
		if err != nil {
			utils.LogError(err)
			http.Error(w, "Unauthorized: No auth-token", http.StatusUnauthorized)
			return
		}
		_, err = jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			utils.LogError(err)
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func GenerateJWT(username string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
    	jwt.MapClaims{ 
    	"username": username, 
    	"exp": time.Now().Add(time.Hour * 24).Unix(), 
    })
    tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET"))) //Sign the token with the secret key
    if err != nil {
    	utils.LogError(err)
    	return "", err
    }
    return tokenString, nil
}