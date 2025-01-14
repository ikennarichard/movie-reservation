package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
)

func AddShowtime(c *gin.Context) {

	var showtime models.Showtime
	var movie models.Movie

	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&movie, showtime.MovieID).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Movie not found"})
		return
	}

	if err := config.DB.Create(&showtime).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to save showtime"})
		return
	}

	c.JSON(201, gin.H{"message": "Showtime added successfully", "showtime": showtime})
}

func UpdateShowtime(c *gin.Context) {

	var showtime models.Showtime

	id := c.Param("id")

	if err := config.DB.First(&showtime, id).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Showtime not found"})
	}

	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
	}

	if err := config.DB.Save(&showtime).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update showtime"})
	}

	c.JSON(200, gin.H{"message": "Showtime updated successfully", "showtime": showtime})
}

func DeleteShowtime(c *gin.Context) {

	id := c.Param("id")

	if err := config.DB.Delete(&models.Showtime{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Showtime not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Showtime deleted successfully"})
}
