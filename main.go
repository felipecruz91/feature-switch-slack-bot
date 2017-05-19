package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

// Slack Bot Token
var slackToken = os.Getenv("SLACK_TOKEN")

func main() {
	slackAPI := slack.New(slackToken)
	rtm := slackAPI.NewRTM()

	// Connect to the Slack RTM API
	go rtm.ManageConnection()

Loop:

	for {
		select {
		case msg := <-rtm.IncomingEvents:

			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Printf("Connection counter: ", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message :%v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> check", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {

					// Parse Experiment name
					experimentName, err := ParseExperimentName(ev.Text)
					if err != nil {
						log.Print(err)
					}

					// Get Experiment from Experimentation API
					var txt string
					experiment, err := GetExperiment(experimentName)
					if err != nil {
						txt = "The experiment *" + experimentName + "* is currently *switched OFF*  in " + "*QA*" + ":x:"
						log.Print(err)
					} else {
						txt = "The experiment *" + experiment.Name + "* is currently *switched ON*  in " + "*QA*" + ":white_check_mark:"
					}

					// Send response to the Slack channel
					rtm.SendMessage(rtm.NewOutgoingMessage(txt, ev.Channel))
				}

			case *slack.RTMError:
				fmt.Printf("Error :%s\n", ev.Error)

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
				//Take no action
			}
		}
	}
}
