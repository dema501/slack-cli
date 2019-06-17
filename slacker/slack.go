package slacker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moul/http2curl"
	"github.com/pkg/errors"
)

// Slack is struct
type Slack struct {
	URL string
}

// Message is post payload
type Message struct {
	Text      string `json:"text"`
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// Post is post message to slack
func (s *Slack) Post(msg Message, verbose bool, timeout time.Duration) (err error) {
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", s.URL, buf)

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/532.5 (KHTML, like Gecko) Chrome/4.0.249.0 Safari/532.5")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if verbose {
		if curl, err := http2curl.GetCurlCommand(req); err == nil {
			fmt.Println("[CURL]: ", curl)
		}
	}

	if err != nil {
		return errors.Wrap(err, "Can't do Slack request")
	}

	client := http.Client{Timeout: timeout}

	res, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "Can't do Slack request")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println("[WARN]: ", errors.Wrap(err, "can't close response body"))
		}
	}()

	if res.StatusCode != 200 {
		return errors.New("post.Failed")
	}

	return
}
