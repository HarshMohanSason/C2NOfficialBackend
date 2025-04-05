package middleware

import(
	"os"
	"net/http"
)

func CORSManager(handlerFunc http.Handler) http.Handler {
	/*The ENV variable in the .env has a value setup to detect
	whether headers are required or not. This is because on the 
	vps when this server is deployed, nginx is handling the cors 
	not the go. This will return multiple headers on the 
	production when a request is made
	*/
	return http.HandlerFunc(func(response http.ResponseWriter, receivedRequest *http.Request){	
		if os.Getenv("HEADERS") == "YES"{
			response.Header().Set("Access-Control-Allow-Origin", "*")
			response.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		}
		//Pre flight request 
		if receivedRequest.Method == http.MethodOptions{
			response.WriteHeader(http.StatusOK)
			return
		}

		handlerFunc.ServeHTTP(response, receivedRequest)
	})
}