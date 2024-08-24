package test

import (
	"context"
	"fmt"
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
	ZooKeeper         testcontainers.Container
)

func SetupContainers() {
	ctx := context.Background()

	// Setup PostgreSQL container
	postgresReq := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "iam",
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
		log.Fatalf("Failed to start Redis contaner: %v", err)
	}
	// set up zookeeper
	zooKeeper := testcontainers.ContainerRequest{
		Image: "confluentinc/cp-zookeeper",
		Env: map[string]string{
			"ZOOKEEPER_CLIENT_PORT":         "2181",
			"ZOOKEEPER_ADMIN_ENABLE_SERVER": "false",
		},
		Name:       "zookeeper",
		WaitingFor: wait.ForListeningPort("2181/tcp"),
		Networks:   []string{"test-megastore"},
		Hostname:   "zookeeper",
	}
	ZooKeeper, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: zooKeeper,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start Kafka container: %v", err)
	}

	zookeeperIP, _ := ZooKeeper.ContainerIP(ctx)

	// Setup Kafka container
	kafkaReq := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-kafka",
		ExposedPorts: []string{"9092/tcp"},
		Env: map[string]string{
			"KAFKA_ZOOKEEPER_CONNECT":                        fmt.Sprintf("%s:2181", zookeeperIP),
			"KAFKA_NUM_PARTITIONS":                           "12",
			"KAFKA_COMPRESSION_TYPE":                         "gzip",
			"KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR":         "1",
			"KAFKA_TRANSACTION_STATE_LOG_REPLICATION_FACTOR": "1",
			"KAFKA_TRANSACTION_STATE_LOG_MIN_ISR":            "1",
			"KAFKA_LISTENER_SECURITY_PROTOCOL_MAP":           "PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT",
			"KAFKA_ADVERTISED_LISTENERS":                     "PLAINTEXT://localhost:29092,PLAINTEXT_HOST://localhost:9092",
			"KAFKA_CONFLUENT_SUPPORT_METRICS_ENABLE":         "false",
			"KAFKA_AUTO_CREATE_TOPICS_ENABLE":                "true",
			"KAFKA_AUTHORIZER_CLASS_NAME":                    "kafka.security.authorizer.AclAuthorizer",
			"KAFKA_ALLOW_EVERYONE_IF_NO_ACL_FOUND":           "true",
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
	if err := ZooKeeper.Terminate(context.Background()); err != nil {
		log.Fatalf("Failed to terminate ZooKeeper container: %v", err)
	}
	if err := KafkaContainer.Terminate(context.Background()); err != nil {
		log.Fatalf("Failed to terminate Kafka container: %v", err)
	}
}
