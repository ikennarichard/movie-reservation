package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"user_id"`
	ShowtimeID  uint      `gorm:"not null" json:"showtime_id"`
	Seats      string   `gorm:"not null" json:"reserved_seats"`
	TotalAmount float64 `json:"total_amount"` // cost of reservation
}