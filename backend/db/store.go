package db

import (
	qq "github.com/sneakybueno/qq/app"
	"github.com/sneakybueno/qq/status"
)

// Store temp db
type Store interface {
	CreateEvent(eventRequest *qq.EventRequest) (*qq.Event, *status.Status)
	AllEvents() ([]qq.Event, *status.Status)
	CreateQuestion(questionRequest *qq.QuestionRequest) (*qq.Question, *status.Status)
	QuestionsForEvent(questionRequest *qq.QuestionRequest) ([]*qq.Question, *status.Status)
}
