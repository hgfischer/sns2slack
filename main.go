package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hgfischer/sns2slack/slack"
	"github.com/hgfischer/sns2slack/sns"
	"log/syslog"
	"net/http"
)

var (
	log *syslog.Writer
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
			vars["channel"],
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
		log.Info(fmt.Sprintf("[%s] %s", resp.Status, snsMsg.String()))
	case sns.SubscriptionConfirmation:
		err = snsMsg.ConfirmSubscription()
		if err != nil {
			log.Err(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	var err error
	log, err = syslog.New(syslog.LOG_INFO, "sns2slack")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	router := mux.NewRouter()
	router.HandleFunc(HookHandlerUrl, hookHandler).Methods("POST")
	http.Handle("/", router)

	log.Info("Listening on :8080...")
	http.ListenAndServe(":8080", nil)
}