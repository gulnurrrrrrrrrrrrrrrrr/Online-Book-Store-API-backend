package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
)

var books = make(map[int]models.Book)
var nextBookID = 1

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryIDStr := r.URL.Query().Get("category")
	pageStr := r.URL.Query().Get("page")

	var filteredBooks []models.Book

	if categoryIDStr != "" {
		if catID, err := strconv.Atoi(categoryIDStr); err == nil {
			for _, book := range books {
				if book.CategoryID == catID {
					filteredBooks = append(filteredBooks, book)
				}
			}
		}
	} else {
		for _, book := range books {
			filteredBooks = append(filteredBooks, book)
		}
	}

	page := 1
	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	start := (page - 1) * 5
	end := start + 5
	if start >= len(filteredBooks) {
		json.NewEncoder(w).Encode([]models.Book{})
		return
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	json.NewEncoder(w).Encode(filteredBooks[start:end])
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Title is required and Price must be greater than 0", http.StatusBadRequest)
		return
	}

	book.ID = nextBookID
	nextBookID++
	books[book.ID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	var updatedBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedBook.Title == "" || updatedBook.Price <= 0 {
		http.Error(w, "Title is required and Price must be > 0", http.StatusBadRequest)
		return
	}

	updatedBook.ID = id
	books[id] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	if _, exists := books[id]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id)
	w.WriteHeader(http.StatusNoContent)
}
