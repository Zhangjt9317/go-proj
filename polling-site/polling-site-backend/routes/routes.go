package routes 

import (
	"polling-site/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes (router *gin.Engine){
	// poll related routes 
	router.POST("/polls", handlers.CreatePoll)
	router.GET("/polls", handlers.ListPolls)
	router.POST("/options/:option_id/vote", handlers.Vote)
}