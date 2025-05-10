package service

import (
	"cli/taskmanager/internal/database"
	"cli/taskmanager/internal/model"
	"errors"
	"fmt"
	"time"
)

// Task service manipulating user actions
type TaskService interface {
	AddTask(title, description, priority string) error
	EditTask(id int, title, desc string) error
	DeleteTask(int) error
	CompleteTask(int) error
	ListTasks() []*model.Task
}

type ImplTaskService struct{}

func (ts ImplTaskService) AddTask(title, desc *string, priority *model.Priority) error {
	if *title == "" {
		return errors.New("invalid input title couldn't be empty")
	}

	task := model.Task{
		Title:       *title,
		Description: *desc,
		Priority:    *priority,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	if err := database.Insert(&task); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (ts ImplTaskService) EditTask(id int, title, desc string) error {
	if title == "" {
		return fmt.Errorf(`title can't be blank, please use -edit "id" -title "new_title" -desc "new_desc"`)
	}

	if err := database.Update(id, title, desc); err != nil {
		return fmt.Errorf("couldn't find task by id(%d)", id)
	}

	return nil
}

func (ts ImplTaskService) DeleteTask(id int) error {
	if err := database.Delete(id); err != nil {
		return fmt.Errorf("task not exists by id(%d)", id)
	}

	return nil
}

func (ts ImplTaskService) CompleteTask(id int) error {
	if err := database.Complete(id); err != nil {
		return fmt.Errorf("task not exists by id(%d)", id)
	}

	return nil
}
