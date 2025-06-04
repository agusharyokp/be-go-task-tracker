package models

import (
	"context"
	"errors"
	"github.com/yourname/go-task-tracker/db"
	"github.com/yourname/go-task-tracker/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, firstname, lastname, password) VALUES ($1, $2, $3, $4)"
	_, err := db.Conn.Exec(context.Background(), query, u.Email, u.FirstName, u.LastName, u.Password)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = $1"
	row := db.Conn.QueryRow(context.Background(), query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword) 
	if err != nil {
		return err
	}

	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("invalid password")
	}

	return nil
}