package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hgfischer/sns2slack/slack"
	"github.com/hgfischer/sns2slack/sns"
	"net/http"
)

const (
	HookHandlerUrl = "/hook/{team}/{token}/{channel}/{icon_emoji}/{username}/publish"
)

func hookHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	defer r.Body.Close()
	body := make([]byte, r.ContentLength)
	_, err = r.Body.Read(body)
	if err != nil {
		log.Err(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	snsMsg, err := sns.NewMessageFromJSON(body)
	if err != nil {
		log.Err(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	switch snsMsg.Type {
	case sns.Notification:
		payload := slack.NewPayload(
			"#"+vars["channel"],
			vars["username"],
			vars["icon_emoji"],
			snsMsg.String(),
		)
		client := slack.Client{Team: vars["team"], Token: vars["token"]}
		resp, err := client.Post(payload)
		if err != nil {
			log.Err(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info(fmt.Sprintf(
			"[Slack: %s] [Vars: %#v] [SNS: %#v]",
			resp.Status,
			vars,
			snsMsg,
		))
	case sns.SubscriptionConfirmation:
		err = snsMsg.ConfirmSubscription()
		if err != nil {
			log.Err(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
