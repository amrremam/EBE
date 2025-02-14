package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrremam/EBE/cmd/api"
	"github.com/amrremam/EBE/config"
	"github.com/amrremam/EBE/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", api.RegisterUser)
	router.POST("/login", api.LoginUser)
	router.POST("/tasks", api.CreateTask)
	router.GET("/tasks", api.GetTasks)
	return router
}

func TestRegisterUser(t *testing.T) {
	router := setupRouter()

	user := models.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	login := models.Login{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonLogin, _ := json.Marshal(login)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// Mock JWT Middleware
func addAuthHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	task := models.Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      false,
	}

	jsonTask, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonTask))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
