package data

import (
	"errors"
	"sync"
	"task_manager/models"
	"time"
)

var (
	tasksMu sync.RWMutex
	lastID  = 0
)

func GetAllTask() []models.Task {
	tasksMu.RLock()
	defer tasksMu.RUnlock()

	copyTasks := make([]models.Task, len(models.Tasks))
	copy(copyTasks, models.Tasks)
	return copyTasks
}

func GetTaskByID(id int) (models.Task, error) {
	tasksMu.RLock()
	defer tasksMu.RUnlock()

	for _, task := range models.Tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return models.Task{}, errors.New("task not found")
}

func CreateTask(input models.Task) models.Task {

	tasksMu.Lock()
	defer tasksMu.Unlock()

	lastID++
	input.ID = lastID
	input.CreatedAt = time.Now()
	input.UpdatedAt = input.CreatedAt

	if input.DueDate.IsZero() {
		input.DueDate = input.CreatedAt
	}

	models.Tasks = append(models.Tasks, input)
	return input

}

func UpdateTask(id int, update models.Task) (models.Task, error) {

	tasksMu.Lock()
	defer tasksMu.Unlock()

	for i, task := range models.Tasks {
		if task.ID == id {

			update.ID = id
			update.CreatedAt = task.CreatedAt
			update.UpdatedAt = time.Now()

			if update.DueDate.IsZero() {
				update.DueDate = task.DueDate
			}
			models.Tasks[i] = update
			return models.Tasks[i], nil

		}
	}

	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id int) error {
	tasksMu.Lock()
	defer tasksMu.Unlock()

	for i, task := range models.Tasks {
		if id == task.ID {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}
