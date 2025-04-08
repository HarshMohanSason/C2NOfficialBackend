package models

import (
	"time"
)

type User struct{
	ID 			int      	`json:"id,omitempty"`  
	Name 		string		`json:"name,omitempty"` //Name can be empty as well if the user is trying to login initially 
	Email 		string 		`json:"email"`
	Password 	*string 	`json:"password,omitempty"`  //Nullable password in case user logs in with a third party service
	AuthType	string      `json:"auth_type"` //Google or regular mail
	CreatedAt	time.Time   `json:"created_at"`
	UpdatedAt   time.Time 	`json:"updated_at"`
}