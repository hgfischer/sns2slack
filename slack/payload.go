package slack

type Payload struct {
	Channel   string `json:"channel,omitempty"`
	Username  string `json:"username,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
	Text      string `json:"text,omitempty"`
}

func NewPayload(channel, username, iconEmoji, text string) *Payload {
	return &Payload{
		Channel:   channel,
		Username:  username,
		IconEmoji: iconEmoji,
		Text:      text,
	}
}
