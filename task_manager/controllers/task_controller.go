package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	all := data.GetAllTask()
	c.IndentedJSON(http.StatusOK, gin.H{"tasks": all})
}

func GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid task Id"})
		return
	}

	task, err := data.GetTaskByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Json body"})
		return
	}

	if newTask.Title == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}
	created := data.CreateTask(newTask)
	c.IndentedJSON(http.StatusCreated, created)

}

func UpdateTask(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Task Id"})
		return
	}

	var updatedTask models.Task

	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Json body"})
		return
	}

	updated, err := data.UpdateTask(id, updatedTask)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, updated)

}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid Task ID"})
		return
	}

	if err := data.DeleteTask(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task removed"})

}
