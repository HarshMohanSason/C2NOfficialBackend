package models
	
import (
	"time"
)
type Review struct {
	ID 			uint 			`json:"id,omitempty"`
	Comment 	string 			`json:"comment"`
	Rating		uint8 			`json:"rating"`
	User 		PublicUser 		`json:"user"`
	Images      []string		`json:"images,omitempty"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
}