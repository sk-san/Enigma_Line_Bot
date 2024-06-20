package main

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	// create channel
	// Provide sugesstions
	// Receive a message from a user
	// run backend app based on user's choice
	// send a result
	// Enigma % heroku config:set LINE_BOT_CHANNEL_SECRET=794b943a91d0c90f3680aaa4ff63c690 -a enigma-line-bot
	// secret:794b943a91d0c90f3680aaa4ff63c690
	// heroku config:set LINE_BOT_CHANNEL_TOKEN = +KCBeRQMcJqaJXpe4XMRMx0k1jnd/w1+7/tj6XiLF+nQh6AwGtwOE7GVNQRQUkBW9xvabZOQdmxNUsVq/Oo1ellmLSInMbysC7S2bubJvLEpl8VyTjag7uSxlPanxbEBT2DdnyX1RsUJgLdKILaJ8gdB04t89/1O/w1cDnyilFU= -a go-enigma-line-bot
	// channel +KCBeRQMcJqaJXpe4XMRMx0k1jnd/w1+7/tj6XiLF+nQh6AwGtwOE7GVNQRQUkBW9xvabZOQdmxNUsVq/Oo1ellmLSInMbysC7S2bubJvLEpl8VyTjag7uSxlPanxbEBT2DdnyX1RsUJgLdKILaJ8gdB04t89/1O/w1cDnyilFU=
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage("Hello User")
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
