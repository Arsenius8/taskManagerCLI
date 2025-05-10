package service

import (
	"cli/taskmanager/internal/database"
	"cli/taskmanager/internal/model"
	"fmt"
	"sort"
	"sync"
)

var mu sync.Mutex

// Managing given flags to next action with tasks
func TaskManager(flags model.Flags) error {
	mu.Lock()
	defer mu.Unlock()

	ts := ImplTaskService{}
	dbTasks, dbSelectError := database.Select()

	switch {
	case *flags.Add:
		return ts.AddTask(flags.Title, flags.Desc, flags.Priority)

	case *flags.Edit != -1:
		return ts.EditTask(*flags.Edit, *flags.Title, *flags.Desc)

	case *flags.Complete != -1:
		return ts.CompleteTask(*flags.Complete)

	case *flags.Del != -1:
		return ts.DeleteTask(*flags.Del)

	case *flags.List:
		if dbSelectError != nil {
			return fmt.Errorf("list all tasks error: %w", dbSelectError)
		}

		printTasks(dbTasks)

	case *flags.Filter != "":
		if dbSelectError != nil {
			return fmt.Errorf("list filtered tasks error: %w", dbSelectError)
		}

		var filtered []*model.Task

		if *flags.Filter == "completed" {
			filtered = Filter(dbTasks, func(task *model.Task) bool {
				return task.Completed
			})
		} else if *flags.Filter == "priority" {
			filtered = Filter(dbTasks, func(task *model.Task) bool {
				return task.Priority == *flags.Priority
			})
		}

		printTasks(filtered)
	}

	return nil
}

func printTasks(tasks []*model.Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	for range 100 {
		fmt.Print("=")
	}
	fmt.Println()
	for _, t := range tasks {
		fmt.Printf("ID(%d) %s (%s) [%v] (completed|%t) (%s)\n",
			t.ID, t.Title,
			t.Description,
			model.PriorityToString(t.Priority),
			t.Completed,
			t.CreatedAt.Format("02-01-2006 15:04:05.999"),
		)
	}
	for range 100 {
		fmt.Print("=")
	}
	fmt.Println()
}

// Returns filtered slise of given slise and func how to do filter
func Filter[T any](values []T, process func(T) bool) []T {
	var filtered []T

	for _, v := range values {
		if process(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}
