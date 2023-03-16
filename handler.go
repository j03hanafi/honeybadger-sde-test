package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func server() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", SpamHandler).Methods("POST")

	// handle unknown routes
	router.NotFoundHandler = Handle404()
	return router
}

func SpamHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	rawPayload := json.NewDecoder(r.Body)
	var payload Payload
	err := rawPayload.Decode(&payload)
	if err != nil {
		log.Println(err)
		responseFormatter(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if !isSpamNotification(payload) {
		log.Println("Payload is not a Spam Notification")
		responseFormatter(w, "Payload is not a Spam Notification", http.StatusOK)
		return
	}

	err = pushSlackAlert(payload.Email)
	if err != nil {
		log.Println(err)
		responseFormatter(w, "Failed to send Slack alert", http.StatusInternalServerError)
		return
	}
	log.Println("Alert pushed")
	responseFormatter(w, "Alert pushed", http.StatusOK)
}

func isSpamNotification(payload Payload) bool {
	return payload.Type == "SpamNotification"
}

func pushSlackAlert(email string) error {
	// testing failed with email "notASlackUser@email.com"
	if email == "notASlackUser@email.com" {
		return errors.New("failed to send Slack alert")
	}
	return nil
}

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseFormatter(w, "404 Not Found", http.StatusNotFound)
	})
}

func responseFormatter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, `{"message": "%s"}`, data)
}
