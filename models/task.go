package models

import (
	"context"
	"time"

	"github.com/yourname/go-task-tracker/db"
)

type TaskStatus string
type TaskPriority string

const (
	TaskStatusOpen TaskStatus = "open"
	TaskStatusInprogress TaskStatus = "inprogress"
	TaskStatusCompleted TaskStatus = "completed"
)

const (
	TaskPriorityLow TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh TaskPriority = "high"
)

type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ProjectID int64 `json:"project_id" binding:"required"`
	Priority TaskPriority `json:"priority" binding:"required,oneof=low medium high"`
	Hours int `json:"hours"`
	DueDate *time.Time `json:"due_date"`
	Status TaskStatus `json:"status" binding:"required,oneof=open inprogress completed"`
}

func (t *Task) Save() error {
	query := "INSERT INTO tasks (title, description, project_id, priority, hours, due_date, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	return db.Conn.QueryRow(context.Background(), query, t.Title, t.Description, t.ProjectID, t.Priority, t.Hours, t.DueDate, t.Status).Scan(&t.ID)
}

func (t *Task) Delete() error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := db.Conn.Exec(context.Background(), query, t.ID)
	return err
}

func (t *Task) Update() error {
	query := "UPDATE tasks SET title = $1, description = $2, priority = $3, hours = $4, due_date = $5, status = $6 WHERE id = $7"
	_, err := db.Conn.Exec(context.Background(), query, t.Title, t.Description, t.Priority, t.Hours, t.DueDate, t.Status, t.ID)
	return err
}

func GetTasksByProjectId(projectId int64) ([]Task, error) {
	query := "SELECT id, title, description, project_id, priority, hours, due_date, status FROM tasks WHERE project_id = $1"
	rows, err := db.Conn.Query(context.Background(), query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.ProjectID, &task.Priority, &task.Hours, &task.DueDate, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	query := "SELECT id, title, description, project_id, priority, hours, due_date, status FROM tasks WHERE id = $1"
	row := db.Conn.QueryRow(context.Background(), query, id)
	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.ProjectID, &task.Priority, &task.Hours, &task.DueDate, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func GetTaskByProjectId(projectId int64) (*Task, error) {
	query := "SELECT id, title, description, project_id, priority, hours, due_date, status FROM tasks WHERE project_id = $1"
	row := db.Conn.QueryRow(context.Background(), query, projectId)
	var task Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.ProjectID, &task.Priority, &task.Hours, &task.DueDate, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}