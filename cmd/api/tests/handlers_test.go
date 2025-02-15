package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amrremam/EBE.git/cmd/api"
	"github.com/amrremam/EBE.git/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	config.ConnectDatabase()
	return api.Routes()
}

func TestRegisterUser(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	body := `{"email": "test@example.com", "password": "password123"}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLoginUser(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	body := `{"email": "test@example.com", "password": "password123"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateTask(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	body := `{"title": "Test Task", "description": "Task Description"}`
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTasks(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer test-token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	taskID := "test-task-id"
	body := `{"title": "Updated Task", "description": "Updated Description", "status": true}`
	req, _ := http.NewRequest("PUT", "/tasks/"+taskID, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Bearer test-token")
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTask(t *testing.T) {
	r := setupTestRouter()
	w := httptest.NewRecorder()
	taskID := "test-task-id"
	req, _ := http.NewRequest("DELETE", "/tasks/"+taskID, nil)
	req.Header.Set("Authorization", "Bearer test-token")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
