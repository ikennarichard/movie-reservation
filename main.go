package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	srv := config.LoadDBConfig()
	srv.ConnectDb()

	config.DB.AutoMigrate(
		&models.User{},
		&models.Genre{},
		&models.Movie{},
		&models.Showtime{},
		&models.Reservation{},
	)

	router := gin.Default()
	setupRoutes(router)

	log.Println("Starting server on port 8080...")
	log.Fatal(router.Run(":8080"))
	
}