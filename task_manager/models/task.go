package models

import "time"

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var Tasks = []Task{
	{
		ID:          1,
		Title:       "Build backend API",
		Description: "Implement task management backend in Go",
		Status:      "in-progress",
		DueDate:     time.Now().AddDate(0, 0, 7),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          2,
		Title:       "Design database schema",
		Description: "Plan and create database tables for tasks and users",
		Status:      "todo",
		DueDate:     time.Now().AddDate(0, 0, 3),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          3,
		Title:       "Write API documentation",
		Description: "Document endpoints for task management API using Swagger",
		Status:      "todo",
		DueDate:     time.Now().AddDate(0, 0, 5),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	{
		ID:          4,
		Title:       "Create frontend integration",
		Description: "Connect React frontend to Go backend API",
		Status:      "todo",
		DueDate:     time.Now().AddDate(0, 0, 10),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
}
