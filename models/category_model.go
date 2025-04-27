package models

import (
	"time"
)

type Category struct {
	ID                uint      `json:"id"`
	Name              string    `json:"name"`
	SizeChart         string    `json:"size_chart"`
	CustomizationPDF  string    `json:"customization_pdf"`
	HowToMeasureImage string    `json:"how_to_measure_image"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
type CategorySummary struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
