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

	api := slack.New(os.Getenv("SLACK_OAUTH_TOKEN"))
	preText := "*Hello! Your GitHub Actions build has finished*"
	jenkinsURL := "*Build URL:*" + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	jenkinsBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + jenkinsURL
	preTextField := slack.NewTextBlockObject("mrkdwn", preText + "\n\n", false, false)
	jenkinsBuildDetailsField := slack.NewTextBlockObject("mrkdwn", jenkinsBuildDetails + "\n\n", false, false)
}
