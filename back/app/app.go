package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	InstID string `json:"inst_id"`
}

func HandleApplication(c *gin.Context) {
	var app Application

	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the function to send a Telegram message
	if err := sendTelegramMessage(app); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Application submitted successfully"})
}
