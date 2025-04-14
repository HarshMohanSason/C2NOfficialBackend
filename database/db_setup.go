package database

import (
	"database/sql"
	_ "github.com/lib/pq" 
	"c2nofficialsitebackend/config"
	"sync"
)

//Global instance for the database 
var (
	db *sql.DB
	once sync.Once
)

func ConnectToDB() error {
	var err error
	//Make sure the db connect only runs once.
	once.Do(func() {
		// Connect to database
		db, err = sql.Open("postgres", config.Env.DB_URL)
	})
	config.LogError(err)
	return err
}

func GetDB() *sql.DB {
    return db
}