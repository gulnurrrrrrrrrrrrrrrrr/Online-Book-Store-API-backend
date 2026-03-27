package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
)

var categories = make(map[int]models.Category)
var nextCategoryID = 1

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categoryList []models.Category
	for _, category := range categories {
		categoryList = append(categoryList, category)
	}

	json.NewEncoder(w).Encode(categoryList)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if category.Name == "" {
		http.Error(w, "Category name is required", http.StatusBadRequest)
		return
	}

	category.ID = nextCategoryID
	nextCategoryID++
	categories[category.ID] = category

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
