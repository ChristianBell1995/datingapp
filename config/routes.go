package config

import (
	"github.com/ChristianBell1995/datingapp/api/controllers"
	"github.com/ChristianBell1995/datingapp/api/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// NewGinRouter all the routes are defined here
func newGinRouter(db *gorm.DB) *gin.Engine {

	httpRouter := gin.Default()
	baseController := controllers.BaseController{
		DB: db,
	}

	// User create endpoint
	httpRouter.GET("/user/create", baseController.CreateUser)
	// Profiles endpoint
	httpRouter.GET("/profiles", middleware.ValidateToken(), baseController.ListUsers)
	// User login
	httpRouter.POST("/user/login", baseController.LoginUser)
	// Swipe endpoint
	httpRouter.POST("/swipe", middleware.ValidateToken(), baseController.CreateSwipe)

	return httpRouter
}
