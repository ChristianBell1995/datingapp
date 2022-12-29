package controllers

import (
	"fmt"
	"net/http"

	"github.com/ChristianBell1995/datingapp/api/auth"
	"github.com/ChristianBell1995/datingapp/api/models"
	"github.com/ChristianBell1995/datingapp/api/utils"
	"github.com/gin-gonic/gin"
)

type CreateSwipeResponse struct {
	Match bool
}

// create swipe
func (base *BaseController) CreateSwipe(c *gin.Context) {

	var swipe = models.Swipe{}
	if err := c.ShouldBindJSON(&swipe); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid json provided"})
		return
	}
	swipe.Prepare()

	userId, err := auth.ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if userId != swipe.SwiperID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	if err := swipe.Validate(base.DB); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	savedSwipe, err := swipe.SaveSwipe(base.DB)

	if err != nil {
		formattedError := utils.FormatError(err.Error())

		c.JSON(http.StatusBadRequest, gin.H{"error": formattedError.Error()})
		return
	}

	response := CreateSwipeResponse{
		Match: false,
	}

	if savedSwipe.Preference == "NO" {
		c.JSON(http.StatusOK, gin.H{"data": &response})
		return
	}

	// If there is another saved swipe with the same swiper and swiped then create a match
	swipeBool, err := savedSwipe.IsAMatch(base.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.Match = swipeBool
	c.JSON(http.StatusOK, gin.H{"data": &savedSwipe})
	return
}
