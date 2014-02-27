package sns

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	Notification             = "Notification"
	SubscriptionConfirmation = "SubscriptionConfirmation"
)

type Message struct {
	Type             string
	MessageId        string
	Token            string
	TopicArn         string
	Subject          string
	Message          string
	Timestamp        time.Time
	SignatureVersion string
	Signature        string
	SigningCertURL   string
	SubscribeURL     string
	UnsubscribeURL   string
}

func NewMessageFromJSON(js []byte) (*Message, error) {
	var err error
	snsMsg := &Message{}
	err = json.Unmarshal(js, &snsMsg)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return snsMsg, err
	}

	snsMsg.Timestamp = snsMsg.Timestamp.In(loc)
	return snsMsg, nil
}

func (sm *Message) String() string {
	return fmt.Sprintf("%s [%s] %s", sm.Timestamp.Format(time.RFC3339), sm.Subject, sm.Message)
}

func (sm *Message) ConfirmSubscription() error {
	_, err := http.Get(sm.SubscribeURL)
	if err != nil {
		return err
	}
	return nil
}
