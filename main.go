package main

import (
	"digitalent/task/controllers"
	"digitalent/task/middlewares"
	"digitalent/task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/api")
	protected := r.Group("/api")
	protected.Use(middlewares.JwtAuthMiddleware())

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected.GET("/user", controllers.CurrentUser)
	protected.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
