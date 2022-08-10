package main

import (
	"digitalent/task/controllers"
	"digitalent/task/controllers/task"
	"digitalent/task/middlewares"
	"digitalent/task/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	// User
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	protected.GET("/user", controllers.CurrentUser)

	// Task
	protected.GET("/tasks", task.FindTasks)
	protected.POST("/tasks", task.CreateTask)
	protected.GET("/tasks/:id", task.FindTask)
	protected.PATCH("/tasks/:id", task.UpdateTask)
	protected.PATCH("/tasks/:id/mark", task.MarkDoneTask)
	protected.DELETE("/tasks/:id", task.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
