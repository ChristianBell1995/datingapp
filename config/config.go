// config struct
package config

import (
	"log"

	"github.com/ChristianBell1995/datingapp/api/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer() Server {
	// Load environment variables
	loadEnv()

	// Initialize database
	db := newDatabase()

	// Initialize router
	router := newGinRouter(db)

	// Migrate database
	db.AutoMigrate(&models.User{}, &models.Swipe{})

	return Server{
		DB:     db,
		Router: router,
	}
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
