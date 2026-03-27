package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
)

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

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bookList []models.Book
	for _, book := range books {
		bookList = append(bookList, book)
	}

	json.NewEncoder(w).Encode(bookList)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if book.Title == "" || book.Price <= 0 {
		http.Error(w, "Title is required and Price must be > 0", http.StatusBadRequest)
		return
	}

	book.ID = nextBookID
	nextBookID++

	books[book.ID] = book

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
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