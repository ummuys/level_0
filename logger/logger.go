package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
)

func InitLogger(path string) (*zerolog.Logger, error) {

	//STD-OUT
	file := initLogFile(path)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}

	multiWriter := io.MultiWriter(file, consoleWriter)

	logger := zerolog.New(multiWriter).With().Timestamp().Logger()

	lvlStr := os.Getenv("LOG_LEVEL") // example: "debug", "info", "error"

	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid LOG_LEVEL: %v", err)
	}
	zerolog.SetGlobalLevel(lvl)

	return &logger, nil
}
