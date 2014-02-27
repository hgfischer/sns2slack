package slack

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Team  string
	Token string
}

func (c *Client) url() string {
	return "https://" + c.Team + ".slack.com/services/hooks/incoming-webhook?token=" + c.Token
}

func (c *Client) Post(payload *Payload) (*http.Response, error) {
	js, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader("payload=" + string(js))
	resp, err := http.Post(c.url(), "application/x-www-form-urlencoded", body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return resp, err
}