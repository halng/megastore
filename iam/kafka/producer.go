package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/tanhaok/megastore/logging"
	"go.uber.org/zap"
)

var Producer *kafka.Producer
var NewUserTopic = "notification.active-new-user"

func InitializeKafkaProducer(bootStrapServer string) error {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootStrapServer})
	if err != nil {
		logging.LOGGER.Error("Cannot create new producer. Terminate...", zap.Any("error", err))
		return err
	}
	//defer Producer.Close()
	return nil
}

func PushMessageNewUser(message string) {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &NewUserTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	deliveryChan := make(chan kafka.Event)
	err := Producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		logging.LOGGER.Error("Cannot produce message")
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		logging.LOGGER.Error(fmt.Sprintf("Delivery failed: %v\n", m.TopicPartition.Error))
	} else {
		logging.LOGGER.Info(fmt.Sprintf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset))
	}

	close(deliveryChan)
}
