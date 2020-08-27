package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	

	qq "github.com/sneakybueno/qq/app"
	"github.com/sneakybueno/qq/db"
)

var store *db.SQLStore

func main() {

	var err error
	datab, err := sqlx.Connect("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Fatalf("error starting databases: %s", err)
	} 

	store = db.NewSQLStore(datab)
	store.Init()

	http.HandleFunc("/events", EventsHandler)
	http.HandleFunc("/questions", QuestionHandler)
	http.ListenAndServe(":8080", nil)
}

// EventsHandler handles creating/getting events
func EventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetEvent(w, r)
	case http.MethodPost:
		handleCreateEvent(w, r)
	default:
		http.Error(w, "Only GET and POST are supported", 405)
	}
}

func handleGetEvent(w http.ResponseWriter, r *http.Request) {
	param, _ := r.URL.Query()["event_id"]
	if (len(param) > 0) {
		// get single event
		eventID := strings.Join(param," ")
		handleGetOneEvent(w, r, eventID)
	} else {
		// get all events
		handleGetAllEvents(w, r)
	}
}

func handleGetOneEvent(w http.ResponseWriter, r *http.Request, eventID string) {
	event, status := store.GetEventByID(eventID)

	if status.Error != nil {
		http.Error(w,  fmt.Sprintf("error decoding request: %v", status.Error), 500)
		return
	}
	json.NewEncoder(w).Encode(event)

}

func handleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, status := store.AllEvents()
	if status.Error != nil {
		http.Error(w,  fmt.Sprintf("error decoding request: %v", status.Error), 500)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	eventRequest := &qq.EventRequest{}

	err := json.NewDecoder(r.Body).Decode(eventRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("error decoding request: %v", err), 400)
		return
	}

	event, status := store.CreateEvent(eventRequest)
	if status.Error != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(event)
}

// QuestionHandler handles creating/getting questions
func QuestionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetAllQuestionsForEvent(w, r)
	case http.MethodPost:
		handleCreateQuestion(w, r)
	default:
		http.Error(w, "Only GET and POST are supported", 405)
	}
}

func handleCreateQuestion(w http.ResponseWriter, r *http.Request) {
	questionRequest := &qq.QuestionRequest{}

	err := json.NewDecoder(r.Body).Decode(questionRequest)
	
	if err != nil || len(questionRequest.EventID) == 0 {
		http.Error(w, fmt.Sprintf("error with request: %v", err), 400)
		return
	}

	question, status := store.CreateQuestion(questionRequest)
	if status.Error != nil {
		http.Error(w, status.Error.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(question)
}

func handleGetAllQuestionsForEvent(w http.ResponseWriter, r *http.Request) {
	eventID, ok := r.URL.Query()["event_id"]

	if ok != true || len(eventID) == 0 {
		http.Error(w, fmt.Sprintf("error with request"), 400)
		return
	}
	questions, status := store.QuestionsForEvent(eventID[0])
	if status.Error != nil {
		http.Error(w, status.Error.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(questions)
}