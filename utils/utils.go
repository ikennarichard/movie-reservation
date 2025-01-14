package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
)

var jwtSecret = []byte(os.Getenv("SECRET"))

func GenerateToken(userID uint, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, nil, errors.New("token has expired")
	} 
	return token, claims, nil
	}
	return nil, nil, errors.New("failed to validate token")
}

func CheckRole(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Access denied"})
			return
		}
}

func GetAdminReports(c *gin.Context) {
	var reservations []models.Reservation
	if err := config.DB.Preload("Seats").Preload("Showtime.Movie").Find(&reservations).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to fetch reservations"})
	}

	var totalRevenue int
	var totalReservations int

	for _, reservation := range reservations {
		totalRevenue += len(reservation.Seats) * 3
		totalReservations++
	}

	report := gin.H{
		"total_revenue":      totalRevenue,
		"total_reservations": totalReservations,
		"reservations":       reservations,
	}

	c.JSON(200, report)
}

func GetCurrentUser(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"message": "successfully fetched user",
		"data": user,
	})
}