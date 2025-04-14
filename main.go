package main

import (
	"c2nofficialsitebackend/database" 
	"log"
	"net/http"
	"c2nofficialsitebackend/handlers"
	"c2nofficialsitebackend/middleware"
	"c2nofficialsitebackend/config"
)

func main() {

	//Load the env file
	config.LoadEnv()

	//Initialize logger 
	config.InitLogger()

	//Connect to postgres
	err := database.ConnectToDB()
	if err != nil{
		log.Println("Database connection error: ",err);
	}
	//Close the connection when main is finished
	defer database.GetDB().Close() 
	
	//Routes
	http.Handle("/signup", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignUpFormUserInfo)))
	http.Handle("/signin", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignInFormUserInfo)))
	http.Handle("/returnuser", middleware.CORSManager(middleware.VerifyJWT(http.HandlerFunc(handlers.ReturnUserInfo))))
	
	//Listening at port 8080
	err = http.ListenAndServe(":8080", nil)
	if err != nil {

	}
}