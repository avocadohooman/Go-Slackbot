package main

import (
	"fmt" // for printf

	"os" // accessing environment variables

	"github.com/slack-go/slack" //github.com slack library
)

func main() {
	api := slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C02J0BZHMCK",
		slack.MsgOptionText("Hello this is Hetki's bot speaking!", false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s", channelID, timestamp)
}
