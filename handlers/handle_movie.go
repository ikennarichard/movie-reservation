package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
)

func AddMovie(c *gin.Context) {
	var genres []models.Genre

	var req struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		PosterImage string   `json:"poster_image"`
		Duration int `json:"duration"`
		Genres      []string `json:"genres"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	for _, genreName := range req.Genres {
			var genre models.Genre
			if err := config.DB.FirstOrCreate(&genre, models.Genre{Name: genreName}).Error; err != nil {
					c.AbortWithStatusJSON(500, gin.H{"error": "Failed to process genres"})
					return
			}
			genres = append(genres, genre)
	}

	movie := models.Movie{
		Title: req.Title,
		Description: req.Description,
		Duration: req.Duration,
		PosterImage: req.PosterImage,
		Genres: genres,
	}

	if err := config.DB.Create(&movie).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to save movie to the database", "details": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "Movie added successfully",
		"data":  movie,
	})
}

func UpdateMovie(c *gin.Context) {

	var movie models.Movie
	id := c.Param("id")

	if err := config.DB.First(&movie, id).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&movie).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(200, gin.H{"message": "Movie updated successfully", "movie": movie})
}

func DeleteMovie(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Movie{}, id).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Failed to delete movie"})
		return
	}

	c.JSON(200, gin.H{"message": "Movie deleted successfully"})
}

func GetMovies(c *gin.Context) {
	var movies []models.Movie
	if err := config.DB.Find(&movies).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(200, movies)
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	var movie models.Movie

	if err := config.DB.First(&movie, id).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(200, movie)
}

func GetMoviesByGenre(c *gin.Context) {
	genreName := c.Param("genre")
	fmt.Println(genreName)
	var genreMovies []models.Movie

	if err := config.DB.Preload("Genres").Joins("JOIN movie_genres ON movie_genres.movie_id = movies.id").
	Joins("JOIN genres ON genres.id = movie_genres.genre_id").
	Where("genres.name = ?", genreName).
	Find(&genreMovies).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, genreMovies)
}


