package test

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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

var (
	PostgresContainer testcontainers.Container
	RedisContainer    testcontainers.Container
	KafkaContainer    testcontainers.Container
)

func SetupContainers() {
	ctx := context.Background()

	// Setup PostgreSQL container
	postgresReq := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	var err error
	PostgresContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: postgresReq,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	// Setup Redis container
	redisReq := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForListeningPort("6379/tcp"),
	}
	RedisContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: redisReq,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start Redis container: %v", err)
	}

	// Setup Kafka container
	kafkaReq := testcontainers.ContainerRequest{
		Image:        "wurstmeister/kafka:2.13-2.6.0",
		ExposedPorts: []string{"9092/tcp"},
		Env: map[string]string{
			"KAFKA_ZOOKEEPER_CONNECT":                "zookeeper:2181",
			"KAFKA_ADVERTISED_LISTENERS":             "PLAINTEXT://localhost:9092",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR": "1",
		},
		WaitingFor: wait.ForListeningPort("9092/tcp"),
	}
	KafkaContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: kafkaReq,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start Kafka container: %v", err)
	}
}

func TearDownContainers() {
	if err := PostgresContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("Failed to terminate PostgreSQL container: %v", err)
	}
	if err := RedisContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("Failed to terminate Redis container: %v", err)
	}
	if err := KafkaContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("Failed to terminate Kafka container: %v", err)
	}
}
