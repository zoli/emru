package task

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("test", "body")
	if task.Title != "test" {
		t.Errorf("Expected task title %s, but got %s", "test", task.Title)
	}
	if task.Body != "body" {
		t.Errorf("Expected task body %s, but got %s", "body", task.Body)
	}
	if task.Done {
		t.Errorf("Expected task status %v, but got %v", false, task.Done)
	}
}

func TestString(t *testing.T) {
	task := NewTask("test", "body")
	if task.String() != "Title: test, Body: body" {
		t.Errorf("Expected task string %s, but got %s", "Title: test, Body: body", task.String())
	}
}

func TestToggle(t *testing.T) {
	tests := []struct {
		s      Status
		expect Status
	}{
		{true, false},
		{false, true},
	}
	for i, test := range tests {
		test.s.toggle()
		if test.s != test.expect {
			t.Errorf("Test %d: Expected %s after toggle but got %s", i, test.expect, test.s)
		}
	}
}