package main

import (
	"fmt"

	"github.com/erfan-272758/saber-landing/app"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

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
