package main

import (
	"c2nofficialsitebackend/config"
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/handlers"
	"c2nofficialsitebackend/middleware"
	"log"
	"net/http"
)

func main() {

	//Load the env file
	config.LoadEnv()

	//Initialize logger
	config.InitLogger()

	//Connect to postgres
	err := database.ConnectToDB()
	if err != nil {

		log.Println("Database connection error: ", err)
	}
	//Close the connection when main is finished
	defer database.GetDB().Close()

	//Set up the uploads directory
	err = config.SetupUploadsDir()
	if err != nil {
		log.Fatal("SetupUploadsDir error: ", err)
	}

	//Serve the Uploads Folder
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("../uploads"))))

	//Routes
	http.Handle("/signup", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignUpFormUserInfo)))
	http.Handle("/signin", middleware.CORSManager(http.HandlerFunc(handlers.ReceiveSignInFormUserInfo)))
	http.Handle("/returnuser", middleware.CORSManager(middleware.VerifyJWT(http.HandlerFunc(handlers.ReturnUserInfo))))
	http.Handle("/addproduct", middleware.CORSManager(middleware.VerifyJWT(http.HandlerFunc(handlers.AddProductData))))
	http.Handle("/addcategory", middleware.CORSManager(middleware.VerifyJWT(http.HandlerFunc(handlers.AddCategoryData))))
	http.Handle("/returnallcategorysummary", middleware.CORSManager(middleware.VerifyJWT(http.HandlerFunc(handlers.ReturnAllCategoriesHandler))))
	//Listening at port 8080
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
