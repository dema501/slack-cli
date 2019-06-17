package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/takecy/slack-cli/slacker"
)

func main() {
	webhookPtr := flag.String("webhook", "", "Webhook url (Required)")
	channelPtr := flag.String("channel", "", "Channel (Required)")
	messagePtr := flag.String("message", "", "Message")
	usernamePtr := flag.String("username", "Incoming-Webhook", "Username")
	iconPtr := flag.String("icon", ":ghost:", "Icon")
	verbosePtr := flag.Bool("verbose", false, "Make the operation more talkative")
	timeoutPtr := flag.Duration("timeout", 10*time.Second, "Timeout specifies a time limit for requests made by this Client.")

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *webhookPtr == "" {
		*webhookPtr = os.Getenv("SLACK_CLI_WEBHOOK")
	}

	if *webhookPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *messagePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	slack := &slacker.Slack{URL: *webhookPtr}
	if err := slack.Post(
		slacker.Message{
			Text:      *messagePtr,
			Username:  *usernamePtr,
			Channel:   *channelPtr,
			IconEmoji: *iconPtr,
		},
		*verbosePtr,
		*timeoutPtr,
	); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
