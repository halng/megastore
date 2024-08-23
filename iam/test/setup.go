package test

import (
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/MyStore/db"
	"github.com/tanhaok/MyStore/models"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func ServeRequest(router *gin.Engine, method string, path string, body string) (int, string) {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res, _ := io.ReadAll(w.Body)
	return w.Code, string(res)
}

func StartPostgresDBInDocker() {
	// Start Postgres DB in Docker
	// Empty for now until I familiarize myself with testcontainers
}

func StartDB() {
	// Start DB
	db.ConnectDB()
	models.Initialize()
}

func TestAll(m *testing.M) {
	//StartDB()
	// additional setup

	code := m.Run()

	// additional teardown

	os.Exit(code)
}
