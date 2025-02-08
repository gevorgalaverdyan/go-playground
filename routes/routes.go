package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/event", createEvent)
	server.PUT("/event/:id", updateEvent)
	server.DELETE("/event/:id", deleteEvent)
}