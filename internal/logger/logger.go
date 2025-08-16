package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	config "github.com/ummuys/level_0/internal/config/logger"
)

func InitLogger(path string) (*config.Loggers, error) {

	zerolog.TimestampFunc = func() time.Time { return time.Now().UTC() }
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	zerolog.TimeFieldFormat = time.RFC3339Nano

	cw := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339Nano,
	}

	//STD-OUT
	file := initLogFile(path)

	multiWriter := io.MultiWriter(file, cw)

	baseLog := zerolog.New(multiWriter).With().Timestamp().Logger()

	logLevels, err := config.ParseLogLevels()
	if err != nil {
		return nil, err
	}

	appLog := baseLog.With().Str("component", "app").Logger().Level(logLevels.AppLvl)
	kfkLog := baseLog.With().Str("component", "kfk").Logger().Level(logLevels.KfkLvl)
	srvLog := baseLog.With().Str("component", "srv").Logger().Level(logLevels.SrvLvl)
	cchLog := baseLog.With().Str("component", "cache").Logger().Level(logLevels.CchLvl)
	svcLog := baseLog.With().Str("component", "svc").Logger().Level(logLevels.SrvLvl)
	dbLog := baseLog.With().Str("component", "db").Logger().Level(logLevels.DbLvl)

	return &config.Loggers{
		AppLog: &appLog,
		KfkLog: &kfkLog,
		SrvLog: &srvLog,
		CchLog: &cchLog,
		SvcLog: &svcLog,
		DbLog:  &dbLog,
	}, nil
}
