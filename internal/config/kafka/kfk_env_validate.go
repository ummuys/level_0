package config

import (
	"fmt"
	"os"
	"strings"
)

func ParseKafkaEnv() (*KafkaEnv, error) {

	var sErr []string

	add := func(env string) {
		sErr = append(sErr, fmt.Sprintf("invalid env for %s", env))
	}

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		add("kafka_broker")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		add("kafka_topic")
	}

	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		add("kafka_group")
	}

	dlqTopic := os.Getenv("KAFKA_DLQ_TOPIC")
	if dlqTopic == "" {
		add("kafka_dlq_topic")
	}

	if len(sErr) > 0 {
		return nil, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return &KafkaEnv{
		Broker:   broker,
		Topic:    topic,
		DlqTopic: dlqTopic,
		Group:    group,
	}, nil

}
