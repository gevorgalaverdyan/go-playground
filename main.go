package main

import (
	"net/http"

	"github.com/gevorgalaverdyan/go-playground/db"
	"github.com/gevorgalaverdyan/go-playground/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	server.GET("/events", getEvents)
	server.POST("/event", createEvent)

	server.Run(":9090")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "internal err"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context){
	var newEvent models.Event

	err := ctx.ShouldBindJSON(&newEvent)
	
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":"binding issue"})
		return 
	}

	errMsg := newEvent.Save()
	if errMsg.Message != "" {
		ctx.JSON(http.StatusInternalServerError, errMsg.Message)
		return
	}
	ctx.JSON(http.StatusCreated, newEvent)
}
