package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/config"
	"github.com/ikennarichard/movie-reservation/models"
	"github.com/ikennarichard/movie-reservation/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	var user models.User

	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Missing or invalid token"})
			return
		}

		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		}
		config.DB.First(&user, claims["user_id"])
		if user.ID == 0 {
			c.AbortWithStatus(401)
		}
		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("user", user)
		c.Set("role", claims["role"].(string))
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.JSON(403, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}
		c.Next()
	}
}
