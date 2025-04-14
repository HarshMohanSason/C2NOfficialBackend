package tests

import (
	"testing"
	"c2nofficialsitebackend/models"
	"c2nofficialsitebackend/database"
	"c2nofficialsitebackend/tests"
)

func TestSearchUser(t *testing.T) {
	
	db, err := tests.SetupTestDB(t)
	if err != nil {
		t.Fatalf("Error setting up the db: %v", err)
	}
	repo := &database.PostgresUserRepository{DB: db}

	for _, tc := range getUserTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			user, err := repo.SearchUser(tc.input)
			if tc.expectErr && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Did not expect error, but got: %v", err)
			}
			t.Logf("Result: %+v", user)
		})
	}
}

func TestInsertUser(t *testing.T){

	db, err := tests.SetupTestDB(t)
	if err != nil{
		t.Fatalf("Error occurred setting up the db: %v", err)
	}
	repo := &database.PostgresUserRepository{DB: db}
	for _, tc := range getUserTestCases(){
		t.Run(tc.name, func(t *testing.T){
			err := repo.CreateUser(tc.input)
			if err != nil{
				t.Fatalf("Error occured, please try again: %v", err)
			}
			t.Logf("Data Insert successfully")
		})
	}
}


func getUserTestCases() []struct {
	name      string
	input     *models.User
	expectErr bool
} {
	return []struct {
		name      string
		input     *models.User
		expectErr bool
	}{
		{
			name: "First User",
			input: &models.User{
				Email:    "abc@gmail.com",
				Password: strPtr("aerearewrw"),
				AuthType: "email",
			},
			expectErr: false,
		},
	}
} 

func strPtr(s string) *string {
	return &s
}