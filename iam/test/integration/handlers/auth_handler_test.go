package handlers

import (
	"github.com/stretchr/testify/assert"
	handlers2 "github.com/tanhaok/megastore/handlers"
	"github.com/tanhaok/megastore/test/integration"
	"net/http"
	"os"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	urlPathLogin := "/api/v1/login"
	router := integration.SetUpRouter()

	router.POST(urlPathLogin, handlers2.Login)

	t.Run("Login: when master credential is used", func(t *testing.T) {
		// Act
		validJsonRequest := `{"password": "changeme", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		// Assert
		if code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, code)
		}
		assert.Equal(t, code, http.StatusOK)
		assert.Contains(t, res, "token")
	})
	t.Run("Login: when data invalid to bind json", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{""password": "changeme" "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Equal(t, res, `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`)
	})
	t.Run("Login: when user is not found", func(t *testing.T) {
		// Act
		validJsonRequest := `{"password": "not-found", "username": "not-found"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		assert.Equal(t, code, http.StatusNotFound)
		assert.Equal(t, res, `{"code":404,"error":"Account not found","status":"ERROR"}`)
	})
	t.Run("Login: when account exist and password is not match", func(t *testing.T) {
		// account register successfully in #TestRegister: Register when user is successfully registered
		// check
		jsonLoginRequest := `{"password": "not-match", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)

		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Invalid credentials","status":"ERROR"}`)

	})
}

//
//func TestCreateStaffHandler(t *testing.T) {
//	urlPathRegister := "/api/v1/create-staff"
//	urlPathLogin := "/api/v1/login"
//
//	headers := map[string]string{
//		constants.ApiTokenRequestHeader:  "test-secret",
//		constants.ApiUserIdRequestHeader: "test-user",
//	}
//
//	router := integration.SetUpRouter()
//	router.POST(urlPathLogin, handlers2.Login)
//	router.POST(urlPathRegister, middleware.ValidateRequest, handlers2.CreateStaff)
//
//	jsonLoginData := "{\"password\": \"changeme\", \"username\": \"changeme\"}"
//
//	t.Run("Create Staff: when request not fulfill - Missing Header", func(t *testing.T) {
//		// Act
//		code, res := integration.ServeRequest(router, "POST", urlPathRegister, "")
//
//		// Assert
//		if code != http.StatusUnauthorized {
//			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
//		}
//		expectedResponse := `{"code":401,"error":"Missing credentials. X-API-SECRET-TOKEN and X-API-USER-ID are required","status":"ERROR"}`
//		assert.Equal(t, expectedResponse, res)
//		assert.Equal(t, 401, code)
//	})
//
//	t.Run("Create Staff: when requester have permission", func(t *testing.T) {
//		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginData)
//
//		assert.Equal(t, 200, code)
//		assert.Contains(t, res, "token")
//
//		var encodedRes dto.LoginResponse
//		err := json.Unmarshal([]byte(res), &encodedRes)
//		assert.NoError(t, err)
//
//		headers[constants.ApiTokenRequestHeader] = encodedRes.ApiToken
//		headers[constants.ApiUserIdRequestHeader] = encodedRes.ID
//
//		code, res, _ = integration.ServeRequestWithHeader(
//			router,
//			"POST",
//			urlPathRegister,
//			`{"email":"test@gmail.com",
//					"username": "test",
//					"password": "test",
//					"lastname": "test",
//					"firstname": "test"}`,
//			headers)
//
//		assert.Equal(t, 201, code)
//
//	})
//
//	t.Run("Create Staff: when email is invalid and missing field", func(t *testing.T) {
//		// Act
//		invalidUserInput := `{"email":"this-is-not-valid-email","username": "changeme", "lastname": "changeme"}`
//
//		code, res, _ := integration.ServeRequestWithHeader(router, "POST", urlPathRegister, invalidUserInput, headers)
//
//		if code != http.StatusBadRequest {
//			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
//		}
//
//		expectedResponse := `{"code":400,"error":{"0":"The email field must be a valid email address","1":"The firstname field is required"},"status":"ERROR"}`
//		assert.Equal(t, expectedResponse, res)
//		assert.Equal(t, 400, code)
//	})
//	t.Run("Create Staff: when data is unable to bind to json", func(t *testing.T) {
//
//		// Act
//		invalidUserInput := `{"email":"this-is-not-valid-email","userName": "changeme", "lastname": "changeme"}`
//		code, res := integration.ServeRequest(router, "POST", urlPathRegister, invalidUserInput)
//
//		// Assert
//		if code != http.StatusBadRequest {
//			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
//		}
//
//		expectedResponse := `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`
//		assert.Equal(t, expectedResponse, res)
//		assert.Equal(t, 400, code)
//	})
//	t.Run("Create Staff: when user is successfully registered", func(t *testing.T) {
//		// Act
//		validUserInput := `{"email":"changeme@gmail.com", "username": "changeme", "password": "changeme", "lastname": "changeme", "firstname": "changeme"}`
//		code, res := integration.ServeRequest(router, "POST", urlPathRegister, validUserInput)
//
//		// Assert
//		if code != http.StatusCreated {
//			t.Errorf("Expected status code %d but got %d", http.StatusCreated, code)
//		}
//		expectedResponse := `{"code":201,"data":null,"status":"SUCCESS"}`
//		assert.Equal(t, expectedResponse, res)
//	})
//	t.Run("Create Staff: when user is exists", func(t *testing.T) {
//		// Act
//		validUserInput := `{"email":"changeme@gmail.com", "username": "changeme", "password": "changeme", "lastname": "changeme", "firstname": "changeme"}`
//		code, res := integration.ServeRequest(router, "POST", urlPathRegister, validUserInput)
//
//		// Assert
//		if code != http.StatusBadRequest {
//			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
//		}
//		expectedResponse := `{"code":400,"error":"Account with username: changeme@gmail.com or email: changeme already exists","status":"ERROR"}`
//		assert.Equal(t, expectedResponse, res)
//	})
//}

func TestValidate(t *testing.T) {
	urlPathValidate := "/api/v1/validate"
	router := integration.SetUpRouter()
	router.GET(urlPathValidate, handlers2.Validate)

	t.Run("Validate: when api token is missing", func(t *testing.T) {
		// Act
		code, res := integration.ServeRequest(router, "GET", urlPathValidate, "")

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Unauthorized","status":"ERROR"}`)
	})
	t.Run("Validate: when api token is invalid", func(t *testing.T) {
		// Act
		code, res, _ := integration.ServeRequestWithHeader(router, "GET", urlPathValidate, "", nil)

		// Assert
		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Unauthorized","status":"ERROR"}`)

	})
}

func TestMain(m *testing.M) {
	integration.SetupTestServer()

	code := m.Run()

	integration.TearDownContainers()
	os.Exit(code)
}
