package http

import (
	"github.com/MigueelLago/todo-go/internal/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	taskHandler := handlers.NewTaskHandler(db)

	tasksRoutes := router.Group("/tasks")
	{
		tasksRoutes.GET("/", taskHandler.GetTasks)
		tasksRoutes.POST("/", taskHandler.CreateTask)
		tasksRoutes.GET("/:id", taskHandler.GetTaskByID)
		tasksRoutes.DELETE("/:id", taskHandler.DeleteTask)
	}
	return router
}
