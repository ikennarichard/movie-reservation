package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
	"github.com/ikennarichard/movie-reservation/utils"
	"golang.org/x/crypto/bcrypt"
)


func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "User registration failed"})
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}


func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.First(&user, "email = ?", input.Email).Error; err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24, "", "", false, true)

	c.JSON(200, gin.H{"token": token})
}