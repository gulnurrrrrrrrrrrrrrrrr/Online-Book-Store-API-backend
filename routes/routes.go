package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/auth"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		r.GET("/books", handlers.GetBooks)
		r.POST("/books", handlers.CreateBook)
		r.GET("/books/:id", handlers.GetBookByID)
		r.PUT("/books/:id", handlers.UpdateBook)
		r.DELETE("/books/:id", handlers.DeleteBook)
		r.GET("/books/search", handlers.SearchBooks)

		r.GET("/authors", handlers.GetAuthors)
		r.POST("/authors", handlers.CreateAuthor)

		r.GET("/categories", handlers.GetCategories)
		r.POST("/categories", handlers.CreateCategory)
	}
	return r
}
