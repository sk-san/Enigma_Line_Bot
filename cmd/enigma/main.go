package main

import (
	"Enigma/internal/core"
	"Enigma/pkg/util"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"), os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	isComplete := true
	textToProcess := ""
	initialSetting := ""
	operationChoice := ""

	message := createSuggestionTemplate()
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage && event.ReplyToken != "" {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if !(isComplete) {
						if _, err = bot.ReplyMessage(event.ReplyToken, createSuggestionTemplate()).Do(); err != nil {
							log.Fatal(err)
						}
						isComplete = true
					} else if isComplete && message.Text == "Decryption" || message.Text == "Encryption" {
						operationChoice = message.Text
						msg := linebot.NewTextMessage("Please enter three alphabets as an initial setting and text (e.g abc.Hello Python) Do not forget adding . dot between three alphabets and text")
						if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
							log.Fatal(err)
						}
					} else if isComplete && util.IsValid(message.Text) {
						dotIndex := strings.Index(message.Text, ".")
						initialSetting = message.Text[:dotIndex]
						textToProcess = message.Text[dotIndex+1:]

						msg := linebot.NewTextMessage("Delete your previous setting and textToProcess before encryption or decryption.")
						if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
							log.Fatal(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Enter Valid strings")).Do(); err != nil {
							log.Fatal(err)
						}
					}
				}
			} else {
				result := ""
				e := core.Enigma_machine{}
				if operationChoice == "Encryption" {
					e.SetDefault(initialSetting)
					result = e.Encrypt(textToProcess)
				} else {
					e.SetDefault(initialSetting)
					result = e.Decrypt(textToProcess)
				}
				message := linebot.NewTextMessage(operationChoice + ": " + result + "   Send something if you want to continue")
				if _, err := bot.BroadcastMessage(message).Do(); err != nil {
					log.Fatal(err)
				}

				isComplete = false
				textToProcess = ""
				initialSetting = ""
				operationChoice = ""
			}
		}
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createSuggestionTemplate() *linebot.TemplateMessage {
	buttons := linebot.NewButtonsTemplate(
		"", "What do you want to do?", "Which one？",
		linebot.NewMessageAction("Encryption", "Encryption"),
		linebot.NewMessageAction("Decryption", "Decryption"),
	)
	return linebot.NewTemplateMessage("processing...", buttons)
}
