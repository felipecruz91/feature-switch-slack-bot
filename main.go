package main

import (
	"io"
	"net/http"
	"os"
	"strings"

	"fmt"

	"github.com/nlopes/slack"
)

// Slack Bot Token
var slackToken = os.Getenv("SLACK_TOKEN")

// HealthCheckEndpoint returns 200 OK
func HealthCheckEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	//router := mux.NewRouter()
	//router.HandleFunc("/healthcheck", HealthCheckEndpoint).Methods("GET")
	//log.Fatal(http.ListenAndServe(":12345", router))

	slackAPI := slack.New(slackToken)
	rtm := slackAPI.NewRTM()

	// Connect to the Slack RTM API
	go rtm.ManageConnection()

Loop:

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			fmt.Print("Event Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Printf("Connection counter: ", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message :%v\n", ev)
				info := rtm.GetInfo()
				prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				if ev.User != info.User.ID && strings.HasPrefix(ev.Text, prefix) {

					rtm.SendMessage(rtm.NewOutgoingMessage("What's up buddy!?!?", ev.Channel))
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
