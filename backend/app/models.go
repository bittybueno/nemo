package qq

import (
	"time"
)

type EventRequest struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type Event struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc" db:"desc"`
}

type Question struct {
	ID        string    `json:"id"`
	EventID   string    `json:"event_id" db:"event_id"`
	CreatedAt time.Time `json:"created_timestamp" db:"created_timestamp"`
	Body      string    `json:"body"`
}

type QuestionRequest struct {
	EventID string `json:"event_id"`
	Body    string `json:"body"`
}

type GetQuestion struct {
	EventID string `json:"event_id"`
}

type Comment struct {
	ID         string    `json:"id"`
	EventID    string    `json:"event_id"`
	QuestionID string    `json:"question_id"`
	CreatedAt  time.Time `json:"created_at"`
}
