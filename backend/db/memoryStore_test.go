package db

import (
	"testing"

	qq "github.com/sneakybueno/qq/app"
)

func TestEmptyStoreReturnsNoEvents(t *testing.T) {
	store := &MemoryStore{events: make(map[string]qq.Event), questions: make(map[string][]*qq.Question), comments: make(map[string][]*qq.Comment)}

	events, status := store.AllEvents()
	if len(events) > 0 {
		t.Errorf("Events should be empty, len = %d", len(events))
	}

	if status.Error != nil {
		t.Errorf("Unexpected error: %v",  status.Error)
	}
}
