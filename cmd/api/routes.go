package api


import (
	"github.com/gin-gonic/gin"
	"github.com/amrremam/EBE.git/middleware"
)



func Routes() *gin.Engine {
	r := gin.Default()
	
	// Register routes
	r.POST("/register", RegisterUser)
	// Login routes
	r.POST("/login", LoginUser)

	// Protected routes (Require JWT)
	authRoutes := r.Group("/")
	authRoutes.Use(middleware.JWTMiddleware())
	authRoutes.POST("/tasks", CreateTask)
	authRoutes.GET("/tasks", GetTasks)
	authRoutes.PUT("/tasks/:id", UpdateTask)
	authRoutes.DELETE("/tasks/:id", DeleteTask)

	return r
}