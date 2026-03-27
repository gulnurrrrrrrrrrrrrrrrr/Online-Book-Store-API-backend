package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
)

var (
	books      = make(map[int]models.Book)
	authors    = make(map[int]models.Author)
	categories = make(map[int]models.Category)

	nextBookID     = 1
	nextAuthorID   = 1
	nextCategoryID = 1
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	r.HandleFunc("/authors", getAuthors).Methods("GET")
	r.HandleFunc("/authors", createAuthor).Methods("POST")

	r.HandleFunc("/categories", getCategories).Methods("GET")
	r.HandleFunc("/categories", createCategory).Methods("POST")

	log.Println("Book Store API is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
