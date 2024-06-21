package main

import (
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
					if message.Text == "選択肢を見せて" {
						if _, err = bot.ReplyMessage(event.ReplyToken, createTemplateMessage()).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("「選択肢を見せて」と入力してください")).Do(); err != nil {
							log.Print(err)
						}
					}
				}
			}
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createTemplateMessage() *linebot.TemplateMessage {
	buttons := linebot.NewButtonsTemplate(
		"", "選択してください", "どれにしますか？",
		linebot.NewMessageAction("選択肢1", "選択肢1が選ばれました"),
		linebot.NewMessageAction("選択肢2", "選択肢2が選ばれました"),
		linebot.NewMessageAction("選択肢3", "選択肢3が選ばれました"),
	)
	return linebot.NewTemplateMessage("これはテンプレートメッセージです", buttons)
}
