package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/config"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
	"net/http"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&category)
	c.JSON(http.StatusCreated, category)
}
