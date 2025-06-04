package models

import (
	"context"
	"fmt"
	"github.com/yourname/go-task-tracker/db"
)


type Project struct {
	ID     int64  `json:"id"`
	Title   string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID int64  `json:"user_id"`
}

func (p *Project) Save() error {
	query := "INSERT INTO projects (title, description, user_id) VALUES ($1, $2, $3) RETURNING id"
	err := db.Conn.QueryRow(context.Background(), query, p.Title, p.Description, p.UserID).Scan(&p.ID)
	return err
}

func (p *Project) Delete() error {
	query := "DELETE FROM projects WHERE id = $1"
	fmt.Println(p.ID)
	_, err := db.Conn.Exec(context.Background(), query, p.ID)
	return err
}

func (p *Project) Update() error {
	query := "UPDATE projects SET title = $1, description = $2 WHERE id = $3"
	_, err := db.Conn.Exec(context.Background(), query, p.Title, p.Description, p.ID)
	return err
}

func (p *Project) FindAll(userId int64) ([]Project, error) {
	query := "SELECT id, title, description, user_id FROM projects WHERE user_id = $1"
	rows, err := db.Conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.UserID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func GetProjectById(id int64) (*Project, error) {
	query := "SELECT id, title, description, user_id FROM projects WHERE id = $1"
	row := db.Conn.QueryRow(context.Background(), query, id)
	var project Project
	err := row.Scan(&project.ID, &project.Title, &project.Description, &project.UserID)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

func IsUserAuthorized(userId int64, projectId int64) bool {
	project, err := GetProjectById(projectId)
	if err != nil {
		return false
	}
	return project.UserID == userId
}