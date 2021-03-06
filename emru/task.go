package emru

import (
	"time"
)

type (
	Task struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		Done      Status    `json:"done"`
		CreatedAt time.Time `json:"created_at"`
	}

	// Task status is it done or not
	Status bool
)

func NewTask(title, body string) *Task {
	return &Task{
		Title:     title,
		Body:      body,
		Done:      false,
		CreatedAt: time.Now(),
	}
}

func (s *Status) Toggle() {
	*s = !(*s)
}

func (s *Status) Val() bool {
	return *s == true
}
