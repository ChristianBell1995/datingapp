package controllers

import (
	"net/http"

	"github.com/ChristianBell1995/datingapp/api/auth"
	"github.com/ChristianBell1995/datingapp/api/models"
	"github.com/ChristianBell1995/datingapp/api/utils"
	"github.com/gin-gonic/gin"
)

// Create Users
func (base *BaseController) CreateUser(c *gin.Context) {
	var user = models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided"})
		return
	}
	user.Prepare()

	if err := user.Validate(""); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userCreated, err := user.SaveUser(base.DB)

	if err != nil {
		formattedError := utils.FormatError(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": formattedError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": &userCreated})
	return
}

// List Users
func (base *BaseController) ListUsers(c *gin.Context) {
	user := models.User{}

	userId, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	users, err := user.FindAllOtherUsers(base.DB, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &users})
	return
}

// Login User
func (base *BaseController) LoginUser(c *gin.Context) {
	var user = models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided"})
		return
	}
	user.Prepare()

	if err := user.Validate("login"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := user.ValidateUserPassword(base.DB, user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": &token})
	return
}
