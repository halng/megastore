package handlers

import (
	"github.com/stretchr/testify/assert"
	"github.com/tanhaok/MyStore/logging"
	"github.com/tanhaok/MyStore/test"
	"net/http"
	"os"
	"testing"
)

func TestRegister(t *testing.T) {
	urlPath := "/api/v1/register"

	t.Run("when email is invalid and missing field", func(t *testing.T) {
		router := test.SetUpRouter()
		router.POST(urlPath, Register)

		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","username": "changeme", "password": "changeme", "lastname": "changeme"}`

		code, res := test.ServeRequest(router, "POST", urlPath, invalidUserInput)

		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":{"0":"The email field must be a valid email address","1":"The firstname field is required"},"status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
	t.Run("when data is unable to bind to json", func(t *testing.T) {
		// Arrange
		router := test.SetUpRouter()
		router.POST(urlPath, Register)

		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","userName": "changeme", passWord": "changeme", "lastname": "changeme"}`
		code, res := test.ServeRequest(router, "POST", urlPath, invalidUserInput)

		// Assert
		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
}

func TestMain(m *testing.M) {
	test.SetupContainers()
	logging.InitLogging()

	code := m.Run()
	test.TearDownContainers()
	os.Exit(code)
}
