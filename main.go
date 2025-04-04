package main

import (
	"c2nofficialsitebackend/database" 
	"log"
	"net/http"
	"c2nofficialsitebackend/handlers"
	"c2nofficialsitebackend/utils"
)

func main() {

	//Initialize the connection to Postgres
	db, err := database.ConnectToDB()
	if err != nil{
		log.Println("Database connection error: ",err);
	}
	defer db.Close() //Closing when main is finished

	//Initializing logger to track errors 
	utils.InitLogger()

	http.HandleFunc("/signup", handlers.ReceiveSignUpFormUserInfo)

	err = http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}