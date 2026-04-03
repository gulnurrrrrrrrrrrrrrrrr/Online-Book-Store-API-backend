package handlers

import (
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/config"
	"github.com/gulnurrrrrrrrrrrrrrrrr/Online-Book-Store-API-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAuthors(c *gin.Context) {
	var authors []models.Author
	config.DB.Find(&authors)
	c.JSON(http.StatusOK, authors)
}

func CreateAuthor(c *gin.Context) {
	var author models.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&author)
	c.JSON(http.StatusCreated, author)
}
