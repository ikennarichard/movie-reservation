package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
	"gorm.io/gorm/clause"
)

type ReserveSeatsRequest struct {
	UserID     uint   `json:"user_id" binding:"required"`
	ShowtimeID uint   `json:"showtime_id" binding:"required"`
	Seats      string `json:"seats" binding:"required"`
	TotalAmount     float64 `json:"total_amount"`
}

func GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("id")

	var showtime models.Showtime
	if err := config.DB.First(&showtime, showtimeID).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Showtime not found"})
		return
	}

	var reservedSeats int
	config.DB.Model(&models.Reservation{}).Where("showtime_id = ?", showtimeID).Select("SUM(seats)").Scan(&reservedSeats)

	availableSeats := showtime.AvailableSeats - reservedSeats

	c.JSON(http.StatusOK, gin.H{
		"showtime_id":     showtimeID,
		"available_seats": availableSeats,
	})
}

func GetUserReservations(c *gin.Context) {
	userID := c.GetUint("user_id")

	var reservations []models.Reservation
	if err := config.DB.Where("user_id = ?", userID).Preload("Showtime").Preload("Showtime.Movie").Find(&reservations).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	c.JSON(200, gin.H{"reservations": reservations})
}

func CancelReservation(c *gin.Context) {
	reservationID := c.Param("id")

	var reservation models.Reservation
	if err := config.DB.Preload("Seats").First(&reservation, reservationID).Error; err != nil {
		c.JSON(404, gin.H{"error": "Reservation not found"})
		return
	}

	// Verify the user owns the reservation
	userID := c.GetUint("user_id")
	if reservation.UserID != userID {
		c.JSON(403, gin.H{"error": "You are not authorized to cancel this reservation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation canceled successfully"})
}

func GetMoviesByDate(c *gin.Context) {

		dateStr := c.Query("date")
		if dateStr == "" {
			c.JSON(400, gin.H{
				"error": "Date is required in the format YYYY-MM-DD",
			})
			return
		}

		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid date format. Use YYYY-MM-DD.",
			})
			return
		}

		var movies []models.Movie
		err = config.DB.Model(&models.Movie{}).Joins("JOIN showtimes ON showtimes.movie_id = movies.id").Where("DATE(showtimes.start_time) = ?", date).Preload("Showtimes", "DATE(start_time) = ?", date).Find(&movies).
		Error

		if len(movies) == 0 {
			c.JSON(200, gin.H{
				"date":    dateStr,
				"message": "No movies available for the specified date",
			})
			return
		}

		if err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to retrieve movies. Please try again later.",
			})
			return
		}
		c.JSON(200, gin.H{
			"date":    dateStr,
			"movies":  movies,
			"message": "Movies successfully retrieved",
		})
}

func ReserveSeats(c *gin.Context) {
	var req ReserveSeatsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	seats := strings.Split(req.Seats, ",")
	if len(seats) == 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "No seats provided"})
		return
	}

	var showtime models.Showtime
	if err := config.DB.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ?", req.ShowtimeID).
		First(&showtime).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "Showtime not found"})
		return
	}

	if showtime.AvailableSeats < len(seats) {
		c.AbortWithStatusJSON(400, gin.H{"error": "Not enough available seats"})
		return
	}

	price := calculateReservationPrice(len(seats), showtime.Price)

	reservation := models.Reservation{
		UserID:        req.UserID,
		ShowtimeID:    req.ShowtimeID,
		Seats: req.Seats,
		TotalAmount: price,
	}

	if err := config.DB.Create(&reservation).Error; err != nil {
		log.Printf("Error creating reservation: %v", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to reserve seats"})
		return
	}

	showtime.AvailableSeats -= len(seats)
	if err := config.DB.Save(&showtime).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update showtime"})
		return
	}

	c.JSON(201, gin.H{
		"message":   "Seats reserved successfully",
		"reservation": reservation,
	})
}

func calculateReservationPrice(seatCount int, showtimePrice float64) float64 {
	return float64(seatCount) * showtimePrice
}
