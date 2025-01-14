package main

import "github.com/gin-gonic/gin"

func handleReadiness(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Welcome to the Movie Reservation System",
	})
}