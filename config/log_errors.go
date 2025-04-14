package config

import (
	"log"
	"os"
)
//Global pointer to the log file 
var errorLogger *log.Logger 

func InitLogger() {
	file, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open error log file: %v", err)
	}
	// Assign log output to file
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
//Logging errors in the file
func LogError(err error) {
	if err != nil {
		errorLogger.Println(err)
	}
}