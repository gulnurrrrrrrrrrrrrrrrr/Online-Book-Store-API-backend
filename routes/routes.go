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

	protected := r.Group("/api")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/favorites", handlers.GetFavorites)
		protected.POST("/favorites/:bookID", handlers.AddToFavorites)
		protected.DELETE("/favorites/:bookID", handlers.RemoveFromFavorites)

		protected.GET("/books", handlers.GetBooks)
		protected.POST("/books", handlers.CreateBook)
		protected.GET("/books/search", handlers.SearchBooks)
		protected.GET("/books/:id", handlers.GetBookByID)
		protected.PUT("/books/:id", handlers.UpdateBook)
		protected.DELETE("/books/:id", handlers.DeleteBook)

		protected.GET("/authors", handlers.GetAuthors)
		protected.POST("/authors", handlers.CreateAuthor)
		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/categories", handlers.CreateCategory)
	}

	return r
}
