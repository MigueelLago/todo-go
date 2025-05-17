package handlers

import (
	"net/http"

	"github.com/MigueelLago/todo-go/internal/domain/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{
		DB: db,
	}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	var tasks []task.Task
	h.DB.Find(&tasks)
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) GetTaskByID(c *gin.Context) {
	var t task.Task
	id := c.Param("id")
	if err := h.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": t})
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var t task.Task
	if err := ctx.ShouldBindJSON(&t); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Create(&t)
	ctx.JSON(http.StatusCreated, gin.H{"data": t})
}
