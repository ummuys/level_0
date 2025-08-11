package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

func InitLogger(path string) (*zerolog.Logger, *zerolog.Logger, *zerolog.Logger, error) {

	//STD-OUT
	file := initLogFile(path)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}

	multiWriter := io.MultiWriter(file, consoleWriter)

	baseLog := zerolog.New(multiWriter).With().Timestamp().Logger()

	lvlStr := os.Getenv("LOG_LEVEL") // example: "debug", "info", "error"

	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid LOG_LEVEL: %v", err)
	}
	zerolog.SetGlobalLevel(lvl)

	kfkLog := baseLog.With().Str("component", "kafka").Logger()
	srvLog := baseLog.With().Str("component", "server").Logger()

	return &baseLog, &kfkLog, &srvLog, nil
}
