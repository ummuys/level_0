package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/ummuys/level_0/internal/validation"
)

func InitLogger(path string) (*zerolog.Logger, *zerolog.Logger, *zerolog.Logger, *zerolog.Logger, error) {

	//STD-OUT
	file := initLogFile(path)
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}

	multiWriter := io.MultiWriter(file, consoleWriter)

	baseLog := zerolog.New(multiWriter).With().Timestamp().Logger()

	appLvl, srvlvl, kfkLvl, cchLvl, err := validation.ParseLogLevels()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	appLog := baseLog.With().Str("component", "app").Logger().Level(appLvl)
	kfkLog := baseLog.With().Str("component", "kafka").Logger().Level(kfkLvl)
	srvLog := baseLog.With().Str("component", "server").Logger().Level(srvlvl)
	cchLog := baseLog.With().Str("component", "cache").Logger().Level(cchLvl)

	return &appLog, &kfkLog, &srvLog, &cchLog, nil
}
