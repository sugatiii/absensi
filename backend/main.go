package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/sugatiii/absensi/config"
	"github.com/sugatiii/absensi/database"
	"github.com/sugatiii/absensi/handlers"
	"github.com/sugatiii/absensi/middleware"
)

func main() {
	cfg := config.Load()
	database.Connect(cfg.DatabaseURL)
	database.Migrate()

	r := gin.Default()
	r.POST("/api/auth/register", handlers.Register)
	r.POST("/api/auth/login", handlers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/dashboard/summary", handlers.DashboardSummary)
		auth.POST("/absensi/checkin", handlers.CheckIn)
		auth.POST("/absensi/checkout", handlers.CheckOut)
		auth.GET("/absensi/history", handlers.AbsensiHistory)
		auth.GET("/users", handlers.ListUsers)
		auth.POST("/users", handlers.CreateUser)
		auth.PUT("/users/:id", handlers.UpdateUser)
		auth.DELETE("/users/:id", handlers.DeleteUser)
	}

	log.Println("Server running on :8080")
	r.Run(":8080")
}