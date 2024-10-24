package main

import (
	"fmt"
	"time"

	"github.com/erfan-272758/saber-landing/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define the POST route
	router.POST("/api/application", app.HandleApplication)

	// get port
	port, ok := app.ConfGet("PORT")
	if !ok {
		port = "8080"
	}

	// Run the server
	router.Run(fmt.Sprintf(":%s", port))
}
