package api


import (
	"github.com/gin-gonic/gin"
	"github.com/amrremam/EBE.git/middleware"
)



func Routes() *gin.Engine {
	r := gin.Default()
	
	// Register routes
	r.POST("/register", registerUser)
	// Login routes
	r.POST("/login", loginUser)

	// Protected routes (Require JWT)
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.JWTMiddleware())
	authRoutes.POST("/tasks", createTask)
	authRoutes.GET("/tasks", getTasks)
	authRoutes.PUT("/tasks/:id", updateTask)
	authRoutes.DELETE("/tasks/:id", deleteTask)

	return r
}