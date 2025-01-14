package models

import (
	"gorm.io/gorm"
	"time"
)

type Showtime struct {
	gorm.Model
	MovieID  uint   `gorm:"not null" json:"movie_id"`
	StartTime     time.Time `gorm:"not null" json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	AvailableSeats int  `json:"available_seats"`
	Price float64 `gorm:"not null" json:"amount"`
}