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

	isComplete := false
	inputText := ""
	initsetting := ""
	choice := ""

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
						if _, err = bot.ReplyMessage(event.ReplyToken, provideSuggestions()).Do(); err != nil {
							log.Print(err)
						}
						isComplete = true

					} else if isComplete && message.Text == "Decryption" || message.Text == "Encryption" {
						choice = message.Text
						msg := linebot.NewTextMessage("Please enter three alphabets as an initial setting and text (e.g abc.Hello Python) Do not forget adding . dot between three alphabets and text")
						if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
							log.Fatal(err)
						}
					} else if isComplete && util.IsValid(message.Text) {
						dotIndex := strings.Index(message.Text, ".")
						initsetting = message.Text[:dotIndex]
						inputText = message.Text[dotIndex+1:]

						msg := linebot.NewTextMessage("Delete your previous message before encryption or decryption.")
						if _, err := bot.BroadcastMessage(msg).Do(); err != nil {
							log.Fatal(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Enter Valid strings")).Do(); err != nil {
							log.Print(err)
						}
					}
				}
			} else {
				result := ""
				e := core.Enigma_machine{}
				if choice == "Encryption" {
					e.SetDefault(initsetting)
					result = e.Encrypt(inputText)
				} else {
					e.SetDefault(initsetting)
					result = e.Decrypt(inputText)
				}
				message := linebot.NewTextMessage(result)
				if _, err := bot.BroadcastMessage(message).Do(); err != nil {
					log.Fatal(err)
				}

				isComplete = false
				inputText = ""
				initsetting = ""
				choice = ""
			}
		}
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func provideSuggestions() *linebot.TemplateMessage {
	buttons := linebot.NewButtonsTemplate(
		"", "What do you want to do?", "Which one？",
		linebot.NewMessageAction("Encryption", "Encryption"),
		linebot.NewMessageAction("Decryption", "Decryption"),
	)
	return linebot.NewTemplateMessage("choice", buttons)
}
