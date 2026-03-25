package storage

import "github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"

var (
	Books      = make(map[int]models.Book)
	Authors    = make(map[int]models.Author)
	Categories = make(map[int]models.Category)
)
