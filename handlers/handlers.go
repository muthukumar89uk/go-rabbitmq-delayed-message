package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"scheduler/connection"
	"scheduler/helper"
	"scheduler/models"
	"scheduler/producer"
	"time"
)

func SchedulePost(w http.ResponseWriter, r *http.Request) {
	var postReq models.PostRequest

	err := json.NewDecoder(r.Body).Decode(&postReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ch, err := connection.Conn.Channel()
	helper.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	body, err := json.Marshal(postReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scheduleTime, err := time.Parse("2006-01-02 15:04:05", postReq.Delay)
	if err != nil {
		log.Println("Error :", err)
		return
	}

	utcTime := scheduleTime.UTC()
	// log.Println("utcTime ::", utcTime)

	currentTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), 0, time.UTC)

	producer.PublishMessage(ch, "delayed_exchange", "delayed_key", string(body), utcTime.Sub(currentTime).Milliseconds())
	w.WriteHeader(http.StatusAccepted)
}
