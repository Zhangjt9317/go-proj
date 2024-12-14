package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var polls = []Poll{}
var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("polls.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database ")
	}
	db.AutoMigrate(&Poll{}, &Option{})
}

func createPoll(c *gin.Context) {
	var newPoll Poll
	if err := c.ShouldBindJSON(&newPoll); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newPoll.ID = len(polls) + 1
	polls = append(polls, newPoll)
	c.JSON(201, newPoll)
}

func listPolls(c *gin.Context) {
	c.JSON(200, polls)
}

func main() {
	initDB()
	router := gin.Default()
	router.POST("/polls", createPoll)
	router.GET("/polls", listPolls)
	router.Static("/static", "./frontend/dist")
	router.Run(":8080") // Run server on port 8080
}
