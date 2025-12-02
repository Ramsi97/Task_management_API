package models

import "time"

type Task struct {
	MongoID     string    `json:"-" bson:"_id,omitempty"`
	ID          int       `json:"id" bson:"id"`    
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	DueDate     time.Time `json:"due_date" bson:"due_date"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
