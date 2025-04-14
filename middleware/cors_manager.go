package middleware

import(
	"net/http"
	"c2nofficialsitebackend/config"
)
/*
CORSManager conditionally adds CORS headers to the response.

In production, when using a reverse proxy like NGINX (which already sets CORS headers),
we avoid setting them in the Go server to prevent duplicate headers.

Controlled via HEADERS env variable (set to "YES" to enable CORS headers from Go).
*/
func CORSManager(next http.Handler) http.Handler {

	return http.HandlerFunc(func(response http.ResponseWriter, receivedRequest *http.Request){	
		if config.Env.ENV_TYPE == "DEV"{
			response.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			response.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			response.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		//Pre flight request 
		if receivedRequest.Method == http.MethodOptions{
			response.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(response, receivedRequest)
	})
}