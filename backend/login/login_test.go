package login

import (
	"bytes"
	"encoding/json"
	"login_register/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", Login)
	return r
}

func TestLogin_Success(t *testing.T) {
	// Przygotuj dane testowe
	models.User_Registers["john"] = models.User_Register{
		Username: "john",
		Password: "pass123",
		Email:    "john@example.com",
	}

	router := setupRouter()

	payload := map[string]string{
		"username": "john",
		"password": "pass123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

func TestLogin_InvalidCredentials(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{
		"username": "john",
		"password": "wrongpass",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Contains(t, w.Body.String(), "Login failed")
}

func TestLogin_EmptyFields(t *testing.T) {
	router := setupRouter()

	payload := map[string]string{
		"username": "",
		"password": "",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "Please fill all the details")
}

func TestLogin_InvalidJSON(t *testing.T) {
	router := setupRouter()

	// Nieprawidłowy JSON
	body := []byte(`{username: "john"}`)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid request")
}
