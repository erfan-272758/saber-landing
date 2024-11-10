package app

import (
	"net/http"
	"time"

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
	go func() {
		for i := 0; i < 5; i++ {
			if err := sendTelegramMessage(app); err == nil {
				return
			}
			time.Sleep(time.Duration((i+1)*2) * time.Second)
		}
	}()
	// if err := sendTelegramMessage(app); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"status": "Application submitted successfully"})
}
