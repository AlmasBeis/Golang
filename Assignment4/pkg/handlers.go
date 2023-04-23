package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandImage = "image"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	if message.IsCommand() {
		switch message.Command() {
		case commandStart:
			return b.handleStartCommand(message)
		case commandImage:
			return b.handleImage(message)
		default:
			return b.handleUnknownCommand(message)
		}
	}
	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case commandImage:
		return b.handleImage(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Bot started")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command for me ")
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleImage(message *tgbotapi.Message) error {

	// Show typing status to indicate that the bot is fetching an image
	chatID := message.Chat.ID
	action := tgbotapi.NewChatAction(chatID, tgbotapi.ChatTyping)
	_, err := b.bot.Send(action)
	if err != nil {
		return err
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
		return err
	}

	// Parse the response
	var image ImageResponse
	err = json.Unmarshal(resp.Body(), &image)
	if err != nil {
		return err
	}

	// Send the image to the user
	photo := tgbotapi.NewPhotoShare(chatID, image.URLs.Regular)
	_, err = b.bot.Send(photo)
	if err != nil {
		return err
	}

	return nil
}
