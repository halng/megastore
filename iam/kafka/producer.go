package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var Producer *kafka.Producer
var newUserTopic = "notification.active-new-user"

func InitializeKafkaProducer(bootStrapServer string) error {
	var err error
	Producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootStrapServer})
	if err != nil {
		fmt.Println("Cannot create new producer. Terminate...")
		return err
	}
	//defer Producer.Close()
	return nil
}

func PushMessageNewUser(message string) {
	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &newUserTopic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	deliveryChan := make(chan kafka.Event)
	err := Producer.Produce(kafkaMessage, deliveryChan)
	if err != nil {
		fmt.Println("Cannot produce message")
		return
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
