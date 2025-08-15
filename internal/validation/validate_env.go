package validation

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

func ParseLogLevels() (zerolog.Level, zerolog.Level, zerolog.Level, zerolog.Level, error) {
	var sErr []string

	appLvlStr := os.Getenv("LOG_LEVEL_APP")
	appLvl, err := zerolog.ParseLevel(appLvlStr)
	if err != nil || appLvlStr == "" {
		sErr = append(sErr, "invalid level for app")
	}

	srvLvlStr := os.Getenv("LOG_LEVEL_SERVER")
	srvLvl, err := zerolog.ParseLevel(srvLvlStr)
	if err != nil || srvLvlStr == "" {
		sErr = append(sErr, "invalid level for server")
	}

	kfkLvlStr := os.Getenv("LOG_LEVEL_KAFKA")
	kfkLvl, err := zerolog.ParseLevel(kfkLvlStr)
	if err != nil || kfkLvlStr == "" {
		sErr = append(sErr, "invalid level for kafka")
	}

	cchLvlStr := os.Getenv("LOG_LEVEL_CACHE")
	cchLvl, err := zerolog.ParseLevel(cchLvlStr)
	if err != nil || cchLvlStr == "" {
		sErr = append(sErr, "invalid level for cache")
	}

	if len(sErr) > 0 {
		return 0, 0, 0, 0, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return appLvl, srvLvl, kfkLvl, cchLvl, nil
}

func ParseKfkEnv() (string, string, string, string, error) {

	var sErr []string

	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		sErr = append(sErr, "invalid kafka_broker")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		sErr = append(sErr, "invalid kafka_topic")
	}

	group := os.Getenv("KAFKA_GROUP")
	if group == "" {
		sErr = append(sErr, "invalid kafka_group")
	}

	dlqTopic := os.Getenv("KAFKA_DLQ_TOPIC")
	if dlqTopic == "" {
		sErr = append(sErr, "invalid kafka_dql_topic")
	}

	if len(sErr) > 0 {
		return "", "", "", "", fmt.Errorf(strings.Join(sErr, ", "))
	}

	return broker, topic, group, dlqTopic, nil

}

func ParseCchEnv() (int, int, error) {

	var sErr []string

	capStr := os.Getenv("CACHE_CAPACITY")
	cap, err := strconv.Atoi(capStr)
	if err != nil {
		sErr = append(sErr, "invalid cache_capacity")
	}

	ttlOrderSecStr := os.Getenv("TTL_ORDER")
	ttlOrderSec, err := strconv.Atoi(ttlOrderSecStr)
	if err != nil {
		sErr = append(sErr, "invalid ttl_order")
	}

	if len(sErr) > 0 {
		return 0, 0, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return cap, ttlOrderSec, nil
}

func ParseSrvEnv() (string, error) {
	port := os.Getenv("APP_PORT")
	if port == "" {
		return "", fmt.Errorf("invalid app_port")
	}
	return port, nil
}
