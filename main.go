package main

import (
	"c2nofficialsitebackend/database" 
	"log"
	"net/http"
	"c2nofficialsitebackend/handlers"
	"c2nofficialsitebackend/utils"
	"c2nofficialsitebackend/middleware"
)

func main() {

	//Initialize the connection to Postgres
	err := database.ConnectToDB()
	if err != nil{
		log.Println("Database connection error: ",err);
	}
	defer database.GetDB().Close() //Closing when main is finished
	//Initializing logger to track errors 
	utils.InitLogger()

	http.Handle("/signup", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignUpFormUserInfo)))
	http.Handle("/signin", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignInFormUserInfo)))
	
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}