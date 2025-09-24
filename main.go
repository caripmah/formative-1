package main

import (
	"belajar-gin/database"
	"belajar-gin/handlers"
	"belajar-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi DB
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Bioskop{})

	// Router
	r := gin.Default()
	r.POST("/bioskop", handlers.CreateBioskop)

	// Run
	r.Run(":8080")
}
