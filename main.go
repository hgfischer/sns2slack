package main

import (
	"github.com/gorilla/mux"
	"log/syslog"
	"net/http"
)

var (
	log *syslog.Writer
)

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