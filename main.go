package main

import (
	// "net/http"

	// "github.com/gevorgalaverdyan/go-playground/db"
	// "github.com/gevorgalaverdyan/go-playground/routes"
	"fmt"
	"time"

	"github.com/gevorgalaverdyan/go-playground/utils"
	// "github.com/gin-gonic/gin"
)

func main() {
	startTime := time.Now()

	utils.PopulateFile();

	fmt.Println(time.Since(startTime))
	// db.InitDB()
	// server := gin.Default()

	// server.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	// })

	// routes.RegisterRoutes(server)

	// server.Run(":9090")
}


