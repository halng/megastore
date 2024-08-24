package kafka

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKafkaProducer(t *testing.T) {
	// Initialize Kafka Producer
	err := InitializeKafkaProducer("localhost:9092")
	assert.NoError(t, err, "Failed to initialize Kafka producer")
}
