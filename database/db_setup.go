package database

import (
	"database/sql"
	"os"
	_ "github.com/lib/pq" 
	"github.com/joho/godotenv"
	"c2nofficialsitebackend/utils"
	"sync"
)

//Global instance for the database 
var (
	db *sql.DB
	once sync.Once
)

func ConnectToDB() (*sql.DB, error) {
	var err error
	//Make sure the db connect only runs once.
	once.Do(func() {
		// Load .env file
		err = godotenv.Load()
		if err != nil {
			utils.LogError(err)
			return
		}
		// Get database URL
		connStr := os.Getenv("DATABASE_URL")
		if connStr == "" {
			utils.LogError(err)
			return
		}
		// Connect to database
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			utils.LogError(err)
			return
		}
	})
	return db, err
}

func GetDB() *sql.DB {
    return db
}