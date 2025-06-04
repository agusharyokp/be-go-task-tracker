package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect() error {
	var err error
	connStr := "postgres://root:root@localhost:15432/task-tracker?sslmode=disable"
	Conn, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	fmt.Println("Database connected")
	return nil
}
