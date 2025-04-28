package handlers

import (
	"c2nofficialsitebackend/database"
	"encoding/json"
	"net/http"
)

func ReturnAllCategoriesHandler(response http.ResponseWriter, receivedRequest *http.Request) {

	if receivedRequest.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	repo := database.PostgresCategoryRepository{DB: database.GetDB()}
	categories, err := repo.ReturnAllCategories()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(response).Encode(categories); err != nil {
		http.Error(response, "Failed to encode user", http.StatusInternalServerError)
	}
}
