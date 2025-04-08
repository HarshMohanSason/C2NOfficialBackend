package tests

import (
	"testing"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/services"
)

func TestSearchUser(t *testing.T){

	db, err := SetupDB(t) 
	if err != nil{
		t.Fatalf("Error occured setting up the db: %v", err)
	}
	repo := &database.PostgresUserRepository{DB: db}

	testInput := &models.User{
		Email:    "harshsason2000@gmail.com",
		Password: nil,
		AuthType: "email",
	}
	user, err := repo.SearchUser(testInput)
	if err != nil {
		t.Fatalf("expected no error; got: %v", err)
	}
	t.Logf("User found: %+v", user)
}
