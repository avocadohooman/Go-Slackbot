package main

import (
	"fmt" // for printf

	"os" // accessing environment variables

	"github.com/slack-go/slack" //github.com slack library

	"github.com/joho/godotenv" // for accessing .env variables
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	args := os.Args[1:]
	fmt.Println(args)

	if len(args) != 4 {
		fmt.Printf("Usage: githubURL buildResult buildNumber jobName \n")
		return 
	}
	
	api := slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))
	preText := "*Hello! Your GitHub Actions build has finished*"
	githubURL := "*Build URL:*" + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	githubBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + githubURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false)
	githubBuildDetailsField := slack.NewTextBlockObject("mrkdwn", githubBuildDetails + "\n\n", false, false)

	githubBuildDetailsSection := slack.NewSectionBlock(githubBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		githubBuildDetailsSection,
	)

	_, _, _, err = api.SendMessage(
		"C02J0BZHMC",
		msg,
	)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
}
