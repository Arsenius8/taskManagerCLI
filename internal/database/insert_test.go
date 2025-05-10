package database

import (
	"cli/taskmanager/internal/model"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {

	task := model.Task{
		Title:       "Test",
		Description: "",
		Completed:   false,
		Priority:    0,
		CreatedAt:   time.Now(),
	}

	if err := Insert(&task); err != nil {
		t.Errorf("Inserting %v failed: %v", task, err)
	}

}
