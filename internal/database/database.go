package database

import (
	"cli/taskmanager/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitEnv() error {
	return godotenv.Load()
}

// Initialize connection to database
func Init() error {
	err := InitEnv()

	if err != nil {
		err = errors.Join(err, fmt.Errorf(".env file load error - %v", err))
	}

	db_user, exists := os.LookupEnv("USER_POSTGRES")

	if !exists {
		err = errors.Join(err, fmt.Errorf("couldn't find USER of postgres in .env file"))
	}

	db_password, exists := os.LookupEnv("PASSWORD_POSTGRES")

	if !exists {
		err = errors.Join(err, fmt.Errorf("couldn't find PASSWORD of postgres in .env file"))

	}

	db_host, exists := os.LookupEnv("HOST")

	if !exists {
		err = errors.Join(err, fmt.Errorf("couldn't find HOST of postgres in .env file"))

	}

	db_port, exists := os.LookupEnv("PORT")

	if !exists {
		err = errors.Join(err, fmt.Errorf("couldn't find PORT of postgres in .env file"))

	}

	db_name, exists := os.LookupEnv("DATABASE")

	if !exists {
		err = errors.Join(err, fmt.Errorf("couldn't find DATABASE NAME in .env file"))

	}

	if err != nil {
		return err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		db_user, db_password, db_host, db_port, db_name)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("couldn't connect to database: %w", err)
	}

	DB = db

	return nil
}

// Create Table for tasks if not exists
func CreateTable() error {
	err := executeQuery(`
	CREATE TABLE IF NOT EXISTS tasks(
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		completed BOOLEAN DEFAULT FALSE,
		priority TEXT NOT NULL,
		createdAt TEXT NOT NULL
	)
	`)

	if err != nil {
		return fmt.Errorf("couldn't create table %w", err)
	}

	return nil
}

// Insert new task to database by given Task struct
func Insert(task *model.Task) error {
	return executeQuery(`
	INSERT INTO tasks (title, description, completed, priority, createdat) VALUES ($1, $2, $3, $4, $5)`,
		task.Title, task.Description, task.Completed, model.PriorityToString(task.Priority), task.CreatedAt,
	)
}

// Delete task from database table by ID
func Delete(id int) error {
	return executeQuery(`DELETE FROM tasks WHERE id = $1`, id)
}

// Updating task completed state by ID
func Complete(id int) error {
	return executeQuery(`UPDATE tasks SET completed = true WHERE id = $1`, id)
}

// Updating task title and description by ID
func Update(id int, title, desc string) error {
	return executeQuery(`UPDATE tasks SET title = $1, description = $2 WHERE id = $3`, title, desc, id)
}

// Executing given query with args
func executeQuery(query string, args ...any) error {
	_, err := DB.Exec(query, args...)

	if err != nil {
		return err
	}

	return nil
}

type taskRows struct {
	id        int
	title     string
	desc      string
	priority  string
	completed bool
	createdAt string
}

// Receive rows from tasks Table from database and return Tasks slice
func Select() ([]*model.Task, error) {
	rows, err := DB.Query(`SELECT id, title, description, priority, completed, createdat FROM tasks`)

	if err != nil {
		return nil, fmt.Errorf("couldn't get rows from database: %w", err)
	}

	var Tasks []*model.Task
	var tr taskRows
	for rows.Next() {
		err := rows.Scan(&tr.id, &tr.title, &tr.desc, &tr.priority, &tr.completed, &tr.createdAt)

		if err != nil {
			return nil, fmt.Errorf("error while scanning rows: %w", err)
		}

		err = collectTasks(&Tasks, &tr)

		if err != nil {
			return nil, err
		}
	}

	return Tasks, nil
}

// collecting data to given tasks
func collectTasks(tasks *[]*model.Task, tr *taskRows) error {
	createdAtTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", tr.createdAt)

	if err != nil {
		return fmt.Errorf("parsing time error: %w", err)
	}

	priorityM := model.ParsePriority(tr.priority)

	task := model.Task{
		ID:          tr.id,
		Title:       tr.title,
		Description: tr.desc,
		Priority:    *priorityM,
		Completed:   tr.completed,
		CreatedAt:   createdAtTime,
	}

	*tasks = append(*tasks, &task)

	return nil
}
