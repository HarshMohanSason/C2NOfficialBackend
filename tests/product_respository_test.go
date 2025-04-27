package tests

import (
	"testing"
)

func TestProductRepository(t *testing.T) {
	db, err := SetupTestDB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}
