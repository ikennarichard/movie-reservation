package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ikennarichard/movie-reservation/handlers"
	"github.com/ikennarichard/movie-reservation/middleware"
	"github.com/ikennarichard/movie-reservation/utils"
)

func setupRoutes(r *gin.Engine) {

	public := r.Group("/api/v1") 
	{
		public.GET("/healthz", handleReadiness)
		public.POST("/login", handlers.Login)
		public.POST("/signup", handlers.Signup)
		public.GET("/movies", handlers.GetMovies)
	}

	auth := r.Group("/api/v1")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/reservations", handlers.GetUserReservations)
		auth.POST("/reservations", handlers.ReserveSeats)
		auth.DELETE("/reservations/:id", handlers.CancelReservation)
		auth.GET("/showtimes/:id/seats", handlers.GetAvailableSeats)
	}

	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/movies", handlers.AddMovie)
		admin.PUT("/movies/:id", handlers.UpdateMovie)
		admin.DELETE("/movies/:id", handlers.DeleteMovie)
		admin.GET("/reports", utils.GetAdminReports)
		admin.POST("/showtimes", handlers.AddShowtime)
		admin.PUT("/showtimes/:id", handlers.UpdateShowtime)
		admin.DELETE("/showtimes/:id", handlers.DeleteShowtime)
	}

}
