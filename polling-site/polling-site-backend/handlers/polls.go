package handlers

import (
	"net/http"
	"polling-site/database"
	"polling-site/models"
	
	"github.com/gin-gonic/gin"
)

// CreatePoll handles creating a new poll
func CreatePoll(c *gin.Context){
	var newPoll models.Poll
	if err := c.ShouldBindJSON(&newPoll); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save to the database 
	database.DB.Create(&newPoll)
	c.JSON(http.StatusCreated, newPoll)
}

// ListPolls handles listing all polls 
func ListPolls(c *gin.Context){
	var polls []models.Poll
	database.DB.Preload("Options").Find(&polls)
	c.JSON(http.StatusOK, polls)
}

// vote handles voting for an option
func Vote(c *gin.Context){
	var option models.Option
	if err := database.DB.First(&option, c.Param("option_id")).Error; err!=nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}
	option.Votes++
	database.DB.Save(&option)
	c.JSON(http.StatusOK, option)
}