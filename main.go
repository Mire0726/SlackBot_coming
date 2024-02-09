package main

import (
	// "fmt"
	"os"
	"github.com/slack-go/slack"
)

// apiをグローバル変数として定義
var api *slack.Client

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			HandleMessage(ev)
		}
	}
}
