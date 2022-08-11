package main

import (
	"digitalent/task/controllers"
	"digitalent/task/controllers/task"
	"digitalent/task/middlewares"
	"digitalent/task/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()

	r := gin.Default()
	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"*"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowAllOrigins = true

	r.Use(cors.New(config))
	r.LoadHTMLGlob("static/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	// User
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	protected.GET("/user", controllers.CurrentUser)
	protected.GET("/users", controllers.FindUsers)

	// Task
	protected.GET("/tasks", task.FindTasks)
	protected.POST("/tasks", task.CreateTask)
	protected.GET("/tasks/:id", task.FindTask)
	protected.PATCH("/tasks/:id", task.UpdateTask)
	protected.PATCH("/tasks/:id/mark", task.MarkDoneTask)
	protected.DELETE("/tasks/:id", task.DeleteTask)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
