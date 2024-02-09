package main

import (
	"fmt"
	"github.com/slack-go/slack"
)

func HandleMessage(ev *slack.MessageEvent) {
	if ev.Text == "きました" {
		userID := ev.User
		IncreaseCount(userID)
		responseText := fmt.Sprintf("<@%s> すごい！", userID)
		api.PostMessage(ev.Channel, slack.MsgOptionText(responseText, false))
	}
}
