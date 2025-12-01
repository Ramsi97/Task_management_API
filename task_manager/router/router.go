package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.GET("", controllers.GetTasks)
		tasks.POST("", controllers.CreateTask)
		tasks.GET(":id", controllers.GetTaskByID)
		tasks.PUT(":id", controllers.UpdateTask)
		tasks.DELETE(":id", controllers.DeleteTask)
	}

	return r
}
