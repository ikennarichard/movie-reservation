package main

import (
	"log"
	"time"

	"github.com/ikennarichard/movie-reservation/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func seedData(db *gorm.DB) {
	// Seed Users
	users := []models.User{
		{Name: "Ceci", Email: "admin@example.com", Password: "adminpassword120", Role: "admin"},
	}

	for _, user := range users {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)

		if err := db.FirstOrCreate(&models.User{}, user).Error; err != nil {
			log.Fatalf("Could not seed user: %v", err)
		}
	}

	// Seed Genres
	genres := []models.Genre{
		{Name: "Action"},
		{Name: "Comedy"},
		{Name: "Drama"},
		{Name: "Horror"},
		{Name: "Romance"},
	}

	var genreMap = make(map[string]models.Genre)
	for _, genre := range genres {
		if err := db.FirstOrCreate(&genre, models.Genre{Name: genre.Name}).Error; err != nil {
			log.Fatalf("Could not seed genre: %v", err)
		}
		genreMap[genre.Name] = genre
	}

	movies := []models.Movie{
		{Title: "Avengers: Endgame", Description: "The Avengers must assemble once again to stop Thanos and save the universe.", Duration: 136, PosterImage: "avengers-endgame.jpg", Genres: []models.Genre{genreMap["Action"], genreMap["Drama"]}},
		{Title: "The Hangover", Description: "A group of friends get into a wild adventure after a bachelor party in Las Vegas.", Duration: 136, PosterImage: "the-hangover.jpg", Genres: []models.Genre{genreMap["Comedy"]}},
		{Title: "Titanic", Description: "A love story aboard the ill-fated RMS Titanic.", Duration: 136, PosterImage: "titanic.jpg", Genres: []models.Genre{genreMap["Romance"], genreMap["Drama"]}},
		{Title: "It", Description: "A group of children must face their fears and fight a shape-shifting entity known as 'It'.", Duration: 136, PosterImage: "it.jpg", Genres: []models.Genre{genreMap["Horror"]}},
		{Title: "Spider-Man: No Way Home", Description: "Peter Parker faces challenges that force him to make difficult decisions about his identity.", Duration: 136, PosterImage: "spiderman-no-way-home.jpg", Genres: []models.Genre{genreMap["Action"], genreMap["Comedy"]}},
	}

	for _, movie := range movies {
		if err := db.Create(&movie).Error; err != nil {
			log.Fatalf("Could not seed movie: %v", err)
		}
	}

	var movieList []models.Movie
	if err := db.Find(&movieList).Error; err != nil {
		log.Fatalf("Could not fetch movies for showtimes: %v", err)
	}

	movieMap := make(map[string]uint)
	for _, movie := range movieList {
		movieMap[movie.Title] = movie.ID
	}

	// Seed Showtimes
	showtimes := []models.Showtime{
		{MovieID: movieMap["Avengers: Endgame"], StartTime: time.Date(2023, 12, 25, 15, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 25, 18, 0, 0, 0, time.UTC)},
		{MovieID: movieMap["The Hangover"], StartTime: time.Date(2023, 12, 25, 20, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 25, 22, 0, 0, 0, time.UTC)},
		{MovieID: movieMap["The Hangover"], StartTime: time.Date(2023, 12, 26, 14, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 26, 17, 0, 0, 0, time.UTC)},
		{MovieID: movieMap["Titanic"], StartTime: time.Date(2023, 12, 26, 19, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 26, 21, 30, 0, 0, time.UTC)},
		{MovieID: movieMap["It"], StartTime: time.Date(2023, 12, 27, 16, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 27, 18, 30, 0, 0, time.UTC)},
		{MovieID: movieMap["Spider-Man: No Way Home"], StartTime: time.Date(2023, 12, 28, 19, 0, 0, 0, time.UTC), AvailableSeats: 100, Price: 10, EndTime: time.Date(2023, 12, 28, 20, 30, 0, 0, time.UTC)},
	}

	for _, showtime := range showtimes {
		if err := db.Create(&showtime).Error; err != nil {
			log.Fatalf("Could not seed showtime: %v", err)
		}
	}

	log.Println("Database seeded successfully!")
}


