package handlers

import (
	"github.com/stretchr/testify/assert"
	handlers2 "github.com/tanhaok/MyStore/handlers"
	"github.com/tanhaok/MyStore/test/integration"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	urlPath := "/api/v1/register"

	t.Run("when email is invalid and missing field", func(t *testing.T) {
		router := integration.SetUpRouter()
		router.POST(urlPath, handlers2.Register)

		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","username": "changeme", "password": "changeme", "lastname": "changeme"}`

		code, res := integration.ServeRequest(router, "POST", urlPath, invalidUserInput)

		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":{"0":"The email field must be a valid email address","1":"The firstname field is required"},"status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
	t.Run("when data is unable to bind to json", func(t *testing.T) {
		// Arrange
		router := integration.SetUpRouter()
		router.POST(urlPath, handlers2.Register)

		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","userName": "changeme", passWord": "changeme", "lastname": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPath, invalidUserInput)

		// Assert
		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
}
