package controllers

import (
	"net/http"
	"strconv"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, models.Tasks)
}

func GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if(err != nil) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "Invalid task Id"})
		return 
	}

	for _, task := range models.Tasks{
		if(task.ID == id){
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func CreateTask(c *gin.Context){
	var newTask := models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Json body"})
		return
	}

	newTask.ID = len(models.Tasks) + 1
	newTask.CreatedAt = time.Now()
	newTask.UpdatedAt = time.Now()

	models.Tasks := append(models.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)

}

func UpdateTask(c *gin.Context){

	id, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Task Id"})
		return
	}

	var updatedTask models.Task 

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Json body"})
		return
	}

	for i, task := range models.Tasks {
		if task.ID == id{
			updatedTask.ID = id
			updatedTask.CreatedAt = task.CreatedAt
			updatedTask.UpdatedAt = time.Now()
			models.Tasks[i] = updatedTask
			c.IndentedJSON(http.StatusOK, updatedTask)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(c *gin.Context){
	id, err = strconv.Atoi(c.Params("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	for i, task := range models.Tasks {
		if task.ID == id {
			models.Tasks = append(models.Tasks[:i], models.Tasks[i+1: ]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}