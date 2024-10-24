package app

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Function to send message to Telegram
func sendTelegramMessage(app Application) error {
	message := fmt.Sprintf(
		"<b>Application</b>\n<b>Name:</b> %s\n<b>Phone:</b> %s\n<b>Instagram ID:</b> %s",
		app.Name, app.Phone, app.InstID,
	)

	// Encode the message and prepare the Telegram API URL

	botToken, ok := ConfGet("TEL_BOT_TOKEN")
	if !ok {
		return fmt.Errorf("TEL_BOT_TOKEN not exists")
	}
	chanID, ok := ConfGet("TEL_CHANNEL_ID")
	if !ok {
		return fmt.Errorf("TEL_CHANNEL_ID not exists")
	}

	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)

	body := fmt.Sprintf("{\"chat_id\":\"%s\",\"text\":\"%s\",\"parse_mode\":\"HTML\"}", chanID, message)
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}

	var c *http.Client

	proxyURL, ok := ConfGet("PROXY_URL")
	if ok {
		// has proxy
		pURL, err := url.Parse(proxyURL)
		if err != nil {
			return err
		}
		c = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(pURL),
			},
			Timeout: 10 * time.Second, // Set a timeout for the request
		}
	} else {
		// normal
		c = &http.Client{}
	}

	// Send the POST request to Telegram API
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBodyB, err := io.ReadAll(resp.Body)
	respBody := ""
	if err == nil {
		respBody = string(respBodyB)
	}
	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status: %s, %s", resp.Status, respBody)
	}

	return nil
}
