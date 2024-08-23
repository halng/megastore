package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestRegister_whenInputIncorrect_shouldReturnBadRequest(t *testing.T) {
	// Arrange
	router := setUpRouter()
	router.POST("/api/v1/register", Register)

	// Act
	invalidUserInput := `{"email":"this-is-not-valid-email","username": "changeme", "password": "changeme", "lastname": "changeme"}`
	req, _ := http.NewRequest("POST", "/api/v1/register", strings.NewReader(invalidUserInput))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res, _ := io.ReadAll(w.Body)
	// Assert

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, w.Code)
	}

	expectedResponse := `{"code":400,"error":{"0":"The email field must be a valid email address","1":"The firstname field is required"},"status":"ERROR"}`
	assert.Equal(t, expectedResponse, string(res))
	assert.Equal(t, 400, w.Code)
}

func TestRegister_whenInputNotAbleToBindToJson_shouldReturnBadRequest(t *testing.T) {
	// Arrange
	router := setUpRouter()
	router.POST("/api/v1/register", Register)

	// Act
	invalidUserInput := `{"email":"this-is-not-valid-email","userName": "changeme", passWord": "changeme", "lastname": "changeme"}`
	req, _ := http.NewRequest("POST", "/api/v1/register", strings.NewReader(invalidUserInput))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res, _ := io.ReadAll(w.Body)
	// Assert

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, w.Code)
	}

	expectedResponse := `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`
	assert.Equal(t, expectedResponse, string(res))
	assert.Equal(t, 400, w.Code)
}
