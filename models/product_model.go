package models

import (
	"time"
)

type Product struct {
	ID               uint16    `json:"id"`
	Name             string    `json:"name"`
	CategoryID       uint      `json:"category_id"`
	LongDescription  string    `json:"long_description"`
	ShortDescription string    `json:"short_description"`
	ThumbnailImage   string    `json:"thumbnail_image"`
	CarouselImages   []string  `json:"carousel_images"`
	Slug             string    `json:"slug"`
	Price            uint32    `json:"price"`
	Discount         uint32    `json:"discount"`
	Inventory        uint16    `json:"inventory"`
	SKU              string    `json:"sku"`
	Status           bool      `json:"status"`
	Weight           float64   `json:"weight"`
	Width            float64   `json:"width"`
	Length           float64   `json:"length"`
	Height           float64   `json:"height"`
	Reviews          []Review  `json:"reviews"`
	TotalRating      float64   `json:"total_rating"`
	AmountSold       uint32    `json:"amount_sold"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
