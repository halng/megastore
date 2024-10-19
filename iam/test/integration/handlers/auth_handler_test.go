package handlers

import (
	"github.com/stretchr/testify/assert"
	handlers2 "github.com/tanhaok/megastore/handlers"
	"github.com/tanhaok/megastore/test/integration"
	"net/http"
	"os"
	"testing"
)

func TestAuthHandler(t *testing.T) {
	urlPathRegister := "/api/v1/register"
	urlPathLogin := "/api/v1/login"
	urlPathValidate := "/api/v1/validate"

	router := integration.SetUpRouter()

	router.POST(urlPathLogin, handlers2.Login)
	router.POST(urlPathRegister, handlers2.Register)
	router.GET(urlPathValidate, handlers2.Validate)

	t.Run("Register: when email is invalid and missing field", func(t *testing.T) {
		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","username": "changeme", "password": "changeme", "lastname": "changeme"}`

		code, res := integration.ServeRequest(router, "POST", urlPathRegister, invalidUserInput)

		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":{"0":"The email field must be a valid email address","1":"The firstname field is required"},"status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
	t.Run("Register: when data is unable to bind to json", func(t *testing.T) {

		// Act
		invalidUserInput := `{"email":"this-is-not-valid-email","userName": "changeme", passWord": "changeme", "lastname": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathRegister, invalidUserInput)

		// Assert
		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}

		expectedResponse := `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
		assert.Equal(t, 400, code)
	})
	t.Run("Register: when user is successfully registered", func(t *testing.T) {
		// Act
		validUserInput := `{"email":"changeme@gmail.com", "username": "changeme", "password": "changeme", "lastname": "changeme", "firstname": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathRegister, validUserInput)

		// Assert
		if code != http.StatusCreated {
			t.Errorf("Expected status code %d but got %d", http.StatusCreated, code)
		}
		expectedResponse := `{"code":201,"data":null,"status":"SUCCESS"}`
		assert.Equal(t, expectedResponse, res)
	})
	t.Run("Register: when user is exists", func(t *testing.T) {
		// Act
		validUserInput := `{"email":"changeme@gmail.com", "username": "changeme", "password": "changeme", "lastname": "changeme", "firstname": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathRegister, validUserInput)

		// Assert
		if code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, code)
		}
		expectedResponse := `{"code":400,"error":"Account with username: changeme@gmail.com or email: changeme already exists","status":"ERROR"}`
		assert.Equal(t, expectedResponse, res)
	})

	t.Run("Login: when data invalid to bind json", func(t *testing.T) {
		// Act
		invalidJsonRequest := `{"email":"changeme@email.com" "password": "changeme", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, invalidJsonRequest)

		assert.Equal(t, code, http.StatusBadRequest)
		assert.Equal(t, res, `{"code":400,"error":"Please check your input. Something went wrong","status":"ERROR"}`)
	})
	t.Run("Login: when user is not found", func(t *testing.T) {
		// Act
		validJsonRequest := `{"email":"not-found@email.com", "password": "not-found", "username": "not-found"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, validJsonRequest)

		assert.Equal(t, code, http.StatusNotFound)
		assert.Equal(t, res, `{"code":404,"error":"Account not found","status":"ERROR"}`)
	})
	t.Run("Login: when account exist and password is not match", func(t *testing.T) {
		// account register successfully in #TestRegister: Register when user is successfully registered
		// check
		jsonLoginRequest := `{"email":"changeme@email.com", "password": "not-match", "username": "changeme"}`
		code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)

		assert.Equal(t, code, http.StatusUnauthorized)
		assert.Equal(t, res, `{"code":401,"error":"Invalid credentials","status":"ERROR"}`)

	})
	//t.Run("Login: when account exist and password is match", func(t *testing.T) {
	//	jsonLoginRequest := `{"email":"changeme@email.com", "password": "changeme", "username": "changeme"}`
	//	code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)
	//
	//	assert.Equal(t, code, http.StatusOK)
	//	assert.Contains(t, res, `{"code":200,"data":{"api-token":"`)
	//
	//	var response map[string]interface{}
	//	assert.NoError(t, json.Unmarshal([]byte(res), &response))
	//
	//	assert.Contains(t, response, "data")
	//	assert.Contains(t, response["data"], "api-token")
	//
	//	apiToken := response["data"].(map[string]interface{})["api-token"]
	//	assert.NotEmpty(t, apiToken)
	//
	//	// check cache
	//	hashedMD := utils.ComputeMD5([]string{"changeme"})
	//	cacheKey := hashedMD + "_" + apiToken.(string)
	//	accessToken, err := db.GetDataFromKey(cacheKey)
	//	logging.LOGGER.Info("Error when get data from key", zap.Any("error", err))
	//	//assert.NoError(t, err)
	//	assert.NotEmpty(t, accessToken)
	//})

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
	//t.Run("Validate: when api token is valid", func(t *testing.T) {
	//	// account register successfully in #TestRegister: Register when user is successfully registered
	//
	//	jsonLoginRequest := `{"email":"changeme@email.com", "password": "changeme", "username": "changeme"}`
	//	code, res := integration.ServeRequest(router, "POST", urlPathLogin, jsonLoginRequest)
	//
	//	assert.Equal(t, code, http.StatusOK)
	//
	//	var response map[string]interface{}
	//	assert.NoError(t, json.Unmarshal([]byte(res), &response))
	//	apiToken := response["data"].(map[string]interface{})["api-token"]
	//	assert.NotEmpty(t, apiToken)
	//
	//	// validate
	//	header := map[string]string{constants.ApiTokenRequestHeader: apiToken.(string), constants.ApiUserIdRequestHeader: "changeme"}
	//	code, res, resHeader := integration.ServeRequestWithHeader(router, "GET", urlPathValidate, "", header)
	//
	//	assert.Equal(t, code, http.StatusOK)
	//	assert.Equal(t, res, `{"code":200,"data":null,"status":"SUCCESS"}`)
	//	assert.NotEmpty(t, resHeader.Get(constants.ApiUserIdRequestHeader))
	//	assert.Equal(t, resHeader.Get(constants.ApiUserRoles), "DEFAULT")
	//	assert.Equal(t, resHeader.Get(constants.ApiUserRequestHeader), "changeme")
	//})
}

func TestMain(m *testing.M) {
	integration.SetupTestServer()

	code := m.Run()

	integration.TearDownContainers()
	os.Exit(code)
}
