package db

import (
	"github.com/google/uuid"
	qq "github.com/sneakybueno/qq/app"
	"github.com/sneakybueno/qq/status"
)

// MemoryStore store for events, questions, comments
type MemoryStore struct {
	events    map[string]qq.Event
	questions map[string][]*qq.Question // mapped by event id
	comments  map[string][]*qq.Comment  // mapped by question id
}

// NewMemoryStore creates memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{events: make(map[string]qq.Event), questions: make(map[string][]*qq.Question), comments: make(map[string][]*qq.Comment)}
}

// CreateEvent creates event from request
func (s *MemoryStore) CreateEvent(eventRequest *qq.EventRequest) (*qq.Event,  *status.Status) {
	id := uuid.New().String()
	event := qq.Event{ID: id, Title: eventRequest.Title, Description: eventRequest.Description}
	s.events[id] = event
	s.questions[id] = []*qq.Question{}
	return &event, status.Ok()
}

// AllEvents retrieves all events
func (s *MemoryStore) AllEvents() ([]qq.Event, *status.Status) {
	v := make([]qq.Event, len(s.events))
	for _, value := range s.events {
		v = append(v, value)
	}
	return v, status.Ok()
}

// CreateQuestion creates a question for an event
func (s *MemoryStore) CreateQuestion(questionRequest *qq.QuestionRequest) (*qq.Question, *status.Status) {
	id := uuid.New().String()
	eventID := questionRequest.EventID
	
	if _, ok := s.events[eventID]; !ok {
		return nil, status.NotFound()
	}

	question := qq.Question{ID: id, EventID: eventID, Body: questionRequest.Body}
	s.questions[eventID] = append(s.questions[eventID], &question)

	return &question, status.Ok()
}

// QuestionsForEvent retrieves questions for event
func (s *MemoryStore) QuestionsForEvent(questionRequest *qq.QuestionRequest) ([]*qq.Question, *status.Status) {
	eventID := questionRequest.EventID
	v := make([]*qq.Question, len(s.questions[eventID]))
	for _, value := range s.questions[eventID] {
		v = append(v, value)
	}
	return v, status.Ok()
}
