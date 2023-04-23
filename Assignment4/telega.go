package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
	"strings"
)

// Send any text message to the bot after the bot has been started
const unsplashAccessKey = "Ss9ZFGE9dK0Dng12cLKVU9CdKvUXFpd08qdZzFNziyQ"
const unsplashBaseURL = "https://api.unsplash.com"

type ImageResponse struct {
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}

func main() {
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(os.Getenv("6205463250:AAHd5GJt1XeZG2zsjAEdb5jbOGk-p7T6Xhs"))
	if err != nil {
		log.Fatal(err)
	}

	// Enable debug mode
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create an update channel to receive updates from Telegram
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	// Process incoming updates
	for update := range updates {
		if update.Message == nil { // Ignore any non-message updates
			continue
		}

		// Check if the message text is "image" or "/image"
		if strings.ToLower(update.Message.Text) == "image" || strings.ToLower(update.Message.Text) == "/image" {
			// Show typing status to indicate that the bot is fetching an image
			chatID := update.Message.Chat.ID
			action := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
			_, err := bot.Send(action)
			if err != nil {
				log.Printf("Failed to send chat action: %v", err)
				continue
			}

			// Make API request to Unsplash to get a random image
			headers := map[string]string{
				"Authorization": fmt.Sprintf("Client-ID %s", unsplashAccessKey),
			}
			client := resty.New()
			resp, err := client.R().
				SetHeaders(headers).
				Get(fmt.Sprintf("%s/photos/random?orientation=landscape", unsplashBaseURL))
			if err != nil {
				log.Printf("Failed to fetch image: %v", err)
				continue
			}

			// Parse the response
			var image ImageResponse
			err = json.Unmarshal(resp.Body(), &image)
			if err != nil {
				log.Printf("Failed to parse image response: %v", err)
				continue
			}

			// Send the image to the user
			photo := tgbotapi.NewPhotoShare(chatID, image.URLs.Regular)
			_, err = bot.Send(photo)
			if err != nil {
				log.Printf("Failed to send photo: %v", err)
				continue
			}
		}
	}
}
