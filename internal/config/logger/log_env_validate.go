package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func ParseLogLevels() (*LogLevels, error) {
	var sErr []string

	add := func(env string) {
		sErr = append(sErr, fmt.Sprintf("invalid level for %s", env))
	}

	appLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_APP"))
	if err != nil {
		add("app")
	}

	srvLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_SERVER"))
	if err != nil {
		add("server")
	}

	kfkLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_KAFKA"))
	if err != nil {
		add("kafka")
	}

	dbLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_DATABASE"))
	if err != nil {
		add("database")
	}

	svcLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_SERVICE"))
	if err != nil {
		add("service")
	}

	cchLvl, err := ParseLevel(os.Getenv("LOG_LEVEL_CACHE"))
	if err != nil {
		add("cache")
	}

	if len(sErr) > 0 {
		return nil, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return &LogLevels{
		AppLvl: appLvl,
		SrvLvl: srvLvl,
		SvcLvl: svcLvl,
		DbLvl:  dbLvl,
		KfkLvl: kfkLvl,
		CchLvl: cchLvl,
	}, nil
}

func ParseLevel(levelStr string) (zerolog.Level, error) {
	if levelStr == "" {
		return 0, fmt.Errorf("empty")
	}

	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return 0, err
	}

	return level, nil
}
