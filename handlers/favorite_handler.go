package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/config"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
)

func AddToFavorites(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in token"})
		return
	}

	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	favorite := models.Favorite{
		UserID: userID.(uint),
		BookID: uint(bookID),
	}

	config.DB.Where("user_id = ? AND book_id = ?", userID, bookID).FirstOrCreate(&favorite)

	c.JSON(http.StatusOK, gin.H{"message": "Book added to favorites successfully"})
}

func RemoveFromFavorites(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in token"})
		return
	}

	bookIDStr := c.Param("bookID")
	bookID, err := strconv.Atoi(bookIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	config.DB.Where("user_id = ? AND book_id = ?", userID, bookID).Delete(&models.Favorite{})

	c.JSON(http.StatusOK, gin.H{"message": "Book removed from favorites successfully"})
}

func GetFavorites(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in token"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 50 {
		limit = 10
	}

	offset := (page - 1) * limit

	var favorites []models.Favorite
	var total int64

	config.DB.Model(&models.Favorite{}).Where("user_id = ?", userID).Count(&total)

	config.DB.Preload("Book").Where("user_id = ?", userID).
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&favorites)

	c.JSON(http.StatusOK, gin.H{
		"data":        favorites,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}
