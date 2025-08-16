package config

import "github.com/rs/zerolog"

//---LOGS---

type LogLevels struct {
	AppLvl zerolog.Level
	SrvLvl zerolog.Level
	DbLvl  zerolog.Level
	SvcLvl zerolog.Level
	KfkLvl zerolog.Level
	CchLvl zerolog.Level
}

type Loggers struct {
	AppLog *zerolog.Logger // APP
	SrvLog *zerolog.Logger // SERVER
	DbLog  *zerolog.Logger // DATABASE
	SvcLog *zerolog.Logger // SERVICE
	KfkLog *zerolog.Logger // KAFKA
	CchLog *zerolog.Logger // CACHE
}

//---LOGS---
