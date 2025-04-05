package middleware
/*
import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"io"
)

func generateJWT(username String) (string, error){

	secretKey = os.Getenv("JWT_SECRET") //Storing the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
    	jwt.MapClaims{ 
    	"username": username, 
    	"exp": time.Now().Add(time.Hour * 24).Unix(), 
    })
    tokenString, err := token.SignedString(secretKey) //Sign the token with the secret key
    if err != nil {
    	return "", err
    }
    return "", nil
}
*/