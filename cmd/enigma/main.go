package main

import (
	"Enigma/internal/core"
	"Enigma/pkg/util"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	bot, err := linebot.New(os.Getenv("LINE_BOT_CHANNEL_SECRET"), os.Getenv("LINE_BOT_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	isComplete := false
	tmp_text := ""
	initsetting := ""

	fmt.Println(tmp_text)

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
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if !(isComplete) {
						if _, err = bot.ReplyMessage(event.ReplyToken, provideSuggestions()).Do(); err != nil {
							log.Print(err)
						}
						isComplete = true

					} else if isComplete && util.IsValid(message.Text) {
						initsetting = message.Text
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Enter Valid strings")).Do(); err != nil {
							log.Print(err)
						}
					}
				}
			} else {
				e := core.Enigma_machine{}
				e.SetDefault(initsetting)

				message := linebot.NewTextMessage("This is a test")
				if _, err := bot.BroadcastMessage(message).Do(); err != nil {
					log.Fatal(err)
				}
				isComplete = false
				tmp_text = ""
				initsetting = ""
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
	setinitmsg := "Please set an initial setting(e.g. abc)"
	buttons := linebot.NewButtonsTemplate(
		"", "What do you want to do?", "Which one？",
		linebot.NewMessageAction("Encryption", setinitmsg),
		linebot.NewMessageAction("Decryption", setinitmsg),
	)
	return linebot.NewTemplateMessage("choice", buttons)
}
