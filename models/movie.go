package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	PosterImage string `json:"poster_image"`
	Duration int `json:"duration"` // time in minutes
	Genres []Genre     `gorm:"many2many:movie_genres;"`
}