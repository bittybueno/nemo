package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3" // justify
	"time"

	"fmt"

	"log"
	qq "github.com/sneakybueno/qq/app"
	"github.com/sneakybueno/qq/status"
)

// SQLStore object that points to db
type SQLStore struct {
	db *sqlx.DB
}

// NewSQLStore returns pointer to SQLite Store
func NewSQLStore(db *sqlx.DB) *SQLStore {
	return &SQLStore{db: db}
}

// Init sets up the databases
func (s *SQLStore) Init() {
	var err error

	configSchema := `PRAGMA foreign_keys = ON;`

	_, err = s.db.Exec(configSchema)
	if err != nil {
		log.Fatal(err)
	}

	eventSchema := `CREATE TABLE events (
		id TEXT NOT NULL,
		title TEXT NOT NULL,
		desc TEXT NOT NULL,
		PRIMARY KEY(id)
		);`

	_, err = s.db.Exec(eventSchema)
	if err != nil {
		log.Fatal(err)
	}

	questionSchema := `CREATE TABLE questions (
		id TEXT NOT NULL,
		event_id TEXT NOT NULL,
		body TEXT NOT NULL,
		created_timestamp DATETIME NOT NULL,
		FOREIGN KEY(event_id) REFERENCES events(id)
		);`

	_, err = s.db.Exec(questionSchema)
	if err != nil {
		log.Fatal(err)
	}
}


// CreateEvent adds the event to the db and returns the event
func (s *SQLStore) CreateEvent(eventRequest *qq.EventRequest) (*qq.Event, *status.Status) {
	id := uuid.New().String()
	event := &qq.Event{ID: id, Title: eventRequest.Title, Description: eventRequest.Description}
	_, err := s.db.NamedExec("INSERT INTO events (id, title, desc) VALUES (:id, :title, :desc)", event)
	
	if err != nil {
		return nil, status.NewStatus(500, err)
	}

	return event, status.Ok()
}

// GetEventByID gets info for one event
func (s *SQLStore) GetEventByID(eventID string) (*qq.Event, *status.Status) {
	event := qq.Event{}
	err := s.db.Get(&event, `SELECT * FROM events WHERE id=$1`, eventID)

	if err != nil {
		fmt.Println(err)
		return nil, status.NotFound()
	}
	fmt.Println(event)
	return &event, status.Ok()
}

// AllEvents returns a list of all events
func (s *SQLStore) AllEvents() ([]*qq.Event, *status.Status) {
	var v []*qq.Event
	rows, err := s.db.Queryx("SELECT * FROM events")

	if err != nil {
		return nil, status.NewStatus(500, err)
	}

	// empty store, no error
	if rows == nil {
		return nil, status.Ok()
	}

	for rows.Next() {
		var e qq.Event
		err = rows.StructScan(&e)

		if err == nil {
			v = append(v, &e)
		}
	}
	rows.Close()

	return v, status.Ok()
}

// CreateQuestion inserts a new question into the db
func (s *SQLStore) CreateQuestion(questionRequest *qq.QuestionRequest) (*qq.Question, *status.Status) {
	id := uuid.New().String()
	question := &qq.Question{ID: id, EventID: questionRequest.EventID, CreatedAt: time.Now(), Body: questionRequest.Body}

	_, err := s.db.NamedExec("INSERT INTO questions (id, event_id, created_timestamp, body) VALUES (:id, :event_id, :created_timestamp, :body)", question)
	
	if err != nil {
		return nil, status.NotFound()
	}
	return question, status.Ok()
}

// QuestionsForEvent takes in an event id and returns the associated
// questions for that event
func (s *SQLStore) QuestionsForEvent(eventID string) ([]*qq.Question, *status.Status) {
	var v []*qq.Question

	rows, err := s.db.Queryx(`SELECT * FROM questions WHERE event_id=$1`, eventID)

	if err != nil {
		fmt.Println(err)
		return nil, status.NotFound()
	}

	for rows.Next() {
		var q qq.Question
		err = rows.StructScan(&q)
		v = append(v, &q)
	}
	rows.Close()

	return v, status.Ok()
}

