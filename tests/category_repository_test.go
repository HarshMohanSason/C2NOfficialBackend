package tests

import (
	"c2nofficialsitebackend/database"
	"testing"
)

func TestGetAllCategories(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	repo := database.PostgresCategoryRepository{DB: db}
	categories, err := repo.ReturnAllCategories()
	if err != nil {
		t.Fatal("Failed to return categories:", err)
	}
	if len(categories) != 1 {
		t.Errorf("Expected 1 category, got %d", len(categories))
	}
}
