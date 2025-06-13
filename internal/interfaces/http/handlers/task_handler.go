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

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var t task.Task
	id := c.Param("id")
	if err := h.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}
	h.DB.Delete(&t)
	c.JSON(http.StatusOK, gin.H{"message": "Task deletada com sucesso"})
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

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var t task.Task
	id := c.Param("id")
	if err := h.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.DB.Save(&t)
	c.JSON(http.StatusOK, gin.H{"data": t})
}

func (h *TaskHandler) CompleteTask(c *gin.Context) {
	var t task.Task
	id := c.Param("id")
	if err := h.DB.First(&t, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}

	t.Done = true
	h.DB.Save(&t)
	c.JSON(http.StatusOK, gin.H{"message": "Task marcada como concluída"})
}
