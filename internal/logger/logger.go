package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

func ParseLevel() (zerolog.Level, zerolog.Level, zerolog.Level, zerolog.Level, error) {
	var sErr []string

	appLvlStr, ok := os.LookupEnv("LOG_LEVEL_APP")
	appLvl, err := zerolog.ParseLevel(appLvlStr)
	if err != nil || !ok {
		sErr = append(sErr, "invalid level for app")
	}

	srvLvlStr, ok := os.LookupEnv("LOG_LEVEL_SERVER")
	srvLvl, err := zerolog.ParseLevel(os.Getenv(srvLvlStr))
	if err != nil || !ok {
		sErr = append(sErr, "invalid level for server")
	}

	kfkLvlStr, ok := os.LookupEnv("LOG_LEVEL_KAFKA")
	kfkLvl, err := zerolog.ParseLevel(kfkLvlStr)
	if err != nil || !ok {
		sErr = append(sErr, "invalid level for kafka")
	}

	cchLvlStr, ok := os.LookupEnv("LOG_LEVEL_CACHE")
	cchLvl, err := zerolog.ParseLevel(cchLvlStr)
	if err != nil || !ok {
		sErr = append(sErr, "invalid level for cache")
	}

	if len(sErr) > 0 {
		return 0, 0, 0, 0, fmt.Errorf(strings.Join(sErr, ", "))
	}

	return appLvl, srvLvl, kfkLvl, cchLvl, nil
}

func InitLogger(path string) (*zerolog.Logger, *zerolog.Logger, *zerolog.Logger, *zerolog.Logger, error) {

	//STD-OUT
	file := initLogFile(path)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}

	multiWriter := io.MultiWriter(file, consoleWriter)

	baseLog := zerolog.New(multiWriter).With().Timestamp().Logger()

	appLvl, srvlvl, kfkLvl, cchLvl, err := ParseLevel()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	appLog := baseLog.With().Str("component", "app").Logger().Level(appLvl)
	kfkLog := baseLog.With().Str("component", "kafka").Logger().Level(kfkLvl)
	srvLog := baseLog.With().Str("component", "server").Logger().Level(srvlvl)
	cchLog := baseLog.With().Str("component", "cache").Logger().Level(cchLvl)

	return &appLog, &kfkLog, &srvLog, &cchLog, nil
}
