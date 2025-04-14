package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          	int       		  `json:"id"`
	Name        	string	  		  `json:"name"`
	ImageLink   	sql.NullString    `json:"image_link"`
	AmountSold  	int               `json:"amount_sold"`
	Inventory   	int               `json:"inventory"`
	Category    	string   		  `json:"category"`
	Description 	sql.NullString    `json:"description"`
	Price       	int      		  `json:"price"`
	SKU         	int      		  `json:"sku"`
	Status      	string   		  `json:"status"`
	Slug        	string   		  `json:"slug"`
	Discount    	int      		  `json:"discount"`
	CreatedAt   	time.Time		  `json:"created_at"`
	UpdatedAt   	time.Time		  `json:"updated_at"`
	Color       	string   		  `json:"color"`        
	Weight      	float64  		  `json:"weight"`
}