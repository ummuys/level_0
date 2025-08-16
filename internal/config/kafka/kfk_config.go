package config

type KafkaEnv struct {
	Broker   string
	Topic    string
	DlqTopic string
	Group    string
}
