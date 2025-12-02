package data

import (
	"context"
	"errors"
	"time"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var(
	taskCollection *mongo.Collection
	lastID int)

// Initialize collection reference
func InitTaskCollection(db *mongo.Database) {
	taskCollection = db.Collection("tasks")
}

// GetAllTask retrieves all tasks from MongoDB
func GetAllTask() ([]models.Task, error) {
	var tasks []models.Task

	cursor, err := taskCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetTaskByID retrieves a single task by its ID
func GetTaskByID(id int) (models.Task, error) {
	
	
	var task models.Task
	err := taskCollection.FindOne(context.Background(), bson.M{"id": id}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	return task, nil
}


func CreateTask(input models.Task) (models.Task, error) {
	lastID++ 
	input.ID = lastID

	now := time.Now()
	input.CreatedAt = now
	input.UpdatedAt = now

	if input.DueDate.IsZero() {
		input.DueDate = now
	}

	res, err := taskCollection.InsertOne(context.Background(), input)
	if err != nil {
		return models.Task{}, err
	}

	input.MongoID = res.InsertedID.(primitive.ObjectID).Hex()
	return input, nil
}


func UpdateTask(id int, update models.Task) (models.Task, error) {
	
	
	var existing models.Task
	err := taskCollection.FindOne(context.Background(), bson.M{"id": id}).Decode(&existing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}

	update.ID = id
	update.CreatedAt = existing.CreatedAt
	update.UpdatedAt = time.Now()
	if update.DueDate.IsZero() {
		update.DueDate = existing.DueDate
	}

	_, err = taskCollection.UpdateOne(
		context.Background(),
		bson.M{"id": id},
		bson.M{"$set": update},
	)
	if err != nil {
		return models.Task{}, err
	}

	return update, nil
}

// DeleteTask removes a task by ID
func DeleteTask(id int) error {
	
	res, err := taskCollection.DeleteOne(context.Background(), bson.M{"id": id})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
