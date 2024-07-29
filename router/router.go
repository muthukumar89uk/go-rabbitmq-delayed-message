package router

import (
	"log"
	"net/http"
	"scheduler/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/schedule", handlers.SchedulePost).Methods("POST")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
