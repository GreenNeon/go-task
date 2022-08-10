package task

import (
	"digitalent/task/models"
	"digitalent/task/utils/token"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Preload("Assignee").Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

type CreateTaskInput struct {
	Name     string `json:"name" form:"name"`
	Assignee int64  `json:"assignee_id" form:"assignee_id"`
	Deadline string `json:"deadline" form:"deadline"`
}

func CreateTask(c *gin.Context) {
	// get user
	user_id, err := token.ExtractTokenID(c)
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02"
	deadlineTime, err := time.Parse(layout, input.Deadline)

	if err != nil {
		fmt.Println(err)
	}

	// Create task
	task := models.Task{Name: input.Name, AssigneeID: input.Assignee, Deadline: deadlineTime,
		CreatedByID: int64(user_id), IsDone: false}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"message": "task created successfuly ...", "data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Task
	err := models.DB.Where("id = ?", c.Param("id")).Preload("Assignee").First(&task).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"message": "task updated successfuly ...", "data": task})
}

type MarkInput struct {
	IsDone bool `json:"is_done" form:"is_done"`
}

func MarkDoneTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input MarkInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&task).Updates(map[string]interface{}{"is_done": input.IsDone})

	c.JSON(http.StatusOK, gin.H{"message": "task mark changed successfuly ...", "data": input.IsDone})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"message": "task deleted successfuly ...", "data": task})
}
