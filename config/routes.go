package config

import (
	"github.com/ChristianBell1995/datingapp/api/controllers"
	"github.com/ChristianBell1995/datingapp/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewGinRouter all the routes are defined here
func newGinRouter(db *gorm.DB) *gin.Engine {

	httpRouter := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	httpRouter.Use(cors.New(config))

	baseController := controllers.BaseController{
		DB: db,
	}

	// User create endpoint
	httpRouter.POST("/user/create", baseController.CreateUser)
	// Profiles endpoint
	httpRouter.GET("/profiles", baseController.ListUsers)
	// User login
	httpRouter.POST("/user/login", baseController.LoginUser)
	// Swipe endpoint
	httpRouter.POST("/swipe", middleware.ValidateToken(), baseController.CreateSwipe)

	return httpRouter
}
