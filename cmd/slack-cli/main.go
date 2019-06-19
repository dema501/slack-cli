package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/dema501/slack-cli/pkg/slacker"
)

// isURL check if string is url
func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func main() {
	webhookPtr := flag.String("webhook", os.Getenv("SLACK_CLI_WEBHOOK"), "Webhook url (Required)")
	channelPtr := flag.String("channel", os.Getenv("SLACK_CLI_CHANNEL"), "Channel (Required)")
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

	if isURL(*webhookPtr) == false {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *messagePtr == "" {
		info, err := os.Stdin.Stat()

		if err == nil {
			if info.Mode()&os.ModeNamedPipe != 0 {
				reader := bufio.NewReader(os.Stdin)

				output := strings.Builder{}

				// Grow the buffer to a decent length, so we don't have to continually
				// re-allocate.
				output.Grow(60)

				for {
					input, _, err := reader.ReadRune()
					if err != nil && err == io.EOF {
						break
					}
					output.WriteRune(input)
				}

				*messagePtr = output.String()
			}
		} else {
			fmt.Fprintf(os.Stderr, "[WARN]: Can't read stdin: %v", err)
		}
	}

	if *messagePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	slack := &slacker.Slack{
		URL:     *webhookPtr,
		Verbose: *verbosePtr,
		Timeout: *timeoutPtr,
	}

	if err := slack.Post(
		slacker.Message{
			Text:      *messagePtr,
			Username:  *usernamePtr,
			Channel:   *channelPtr,
			IconEmoji: *iconPtr,
		}); err != nil {
		fmt.Fprintf(os.Stderr, "[ERR]: %v", err)
	}
}
