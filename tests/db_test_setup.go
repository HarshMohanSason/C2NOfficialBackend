package tests

import (
	"database/sql"
	"testing"
)

func SetupDB(t *testing.T) (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=harshmohansason dbname=c2nofficialsitetestdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}