package todo_test

import (
	"testing"
	"todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	taskName := "New Task"
	l.Add(taskName)
	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	taskName := "New Task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("new task should not be completed")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("new task should be completed")
	}
}
