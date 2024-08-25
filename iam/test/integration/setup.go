package integration

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tanhaok/megastore/db"
	kafka2 "github.com/tanhaok/megastore/kafka"
	"github.com/tanhaok/megastore/logging"
	"github.com/tanhaok/megastore/models"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func ServeRequest(router *gin.Engine, method string, path string, body string) (int, string) {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res, _ := io.ReadAll(w.Body)
	return w.Code, string(res)
}

func ServeRequestWithHeader(router *gin.Engine, method string, path string, body string, header map[string]string) (int, string, http.Header) {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	for key, value := range header {
		req.Header.Set(key, value)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res, _ := io.ReadAll(w.Body)
	return w.Code, string(res), w.Header()
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

	//PostgresContainer.Start(ctx)

	// Setup Redis container
	redisReq := testcontainers.ContainerRequest{
		Image:        "redis:6",
		ExposedPorts: []string{"6379/tcp"},
		Env: map[string]string{
			"ALLOW_EMPTY_PASSWORD": "yes",
		},
		WaitingFor: wait.ForListeningPort("6379/tcp"),
	}
	RedisContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: redisReq,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start Redis contaner: %v", err)
	}

	//RedisContainer.Start(ctx)

	// set up zookeeper
	zooKeeper := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-zookeeper",
		ExposedPorts: []string{"2181/tcp"},
		Env: map[string]string{
			"ZOOKEEPER_MAX_RETRIES":         "5",
			"ZOOKEEPER_RETRY_INTERVAL_MS":   "1000",
			"ZOOKEEPER_CLIENT_PORT":         "2181",
			"ZOOKEEPER_ADMIN_ENABLE_SERVER": "false",
		},
		WaitingFor: wait.ForListeningPort("2181/tcp"),
		Networks:   []string{"test-megastore"},
	}
	ZooKeeper, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: zooKeeper,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start Kafka container: %v", err)
	}
	//ZooKeeper.Start(ctx)
	zookeeperIP, _ := ZooKeeper.ContainerIP(ctx)

	// Setup Kafka container
	kafkaReq := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-kafka",
		ExposedPorts: []string{"9092/tcp", "29092/tcp"},
		Env: map[string]string{
			"KAFKA_ZOOKEEPER_CONNECT":                        fmt.Sprintf("%s:2181", zookeeperIP),
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
			"KAFKA_BROKER_ID":                                "0",
		},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort("9092/tcp"),
			wait.ForListeningPort("29092/tcp"),
			wait.ForLog("started (kafka.server.KafkaServer)")),
	}
	KafkaContainer, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: kafkaReq,
		Started:          true,
	})

	if err != nil {
		log.Fatalf("Failed to create Kafka container: %v", err)
		return
	}
	//KafkaContainer.Start(ctx)

}

func TearDownContainers() {
	containers := []testcontainers.Container{
		PostgresContainer,
		RedisContainer,
		ZooKeeper,
		KafkaContainer,
	}
	for _, container := range containers {
		if err := container.Terminate(context.Background()); err != nil {
			log.Fatalf("Failed to terminate container: %v", err)
		}
	}
}

func SetupTestServer() {
	SetupContainers()
	// set up some env variable
	ctx := context.Background()
	//kafkaBootStrapServerIP, _ := KafkaContainer.ContainerIP(ctx)
	kafkaBootStrapServer := fmt.Sprintf("%s:9092", "localhost")
	postgresIP, _ := PostgresContainer.ContainerIP(ctx)
	redisIP, _ := RedisContainer.ContainerIP(ctx)

	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_HOST", postgresIP)
	os.Setenv("DATABASE", "iam")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("REDIS_HOST", redisIP)
	os.Setenv("REDIS_DB", "0")
	os.Setenv("REDIS_PASSWORD", "")

	err := kafka2.InitializeKafkaProducer(kafkaBootStrapServer)
	if err != nil {
		logging.LOGGER.Error("Failed to initialize Kafka producer")
		return
	}

	logging.InitLogging()
	db.ConnectDB()
	models.Initialize()
}
