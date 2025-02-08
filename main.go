package main

import (
	"net/http"

	"github.com/gevorgalaverdyan/go-playground/db"
	"github.com/gevorgalaverdyan/go-playground/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	routes.RegisterRoutes(server)

	server.Run(":9090")
}


