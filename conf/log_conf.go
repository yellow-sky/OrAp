package conf

import (
	"github.com/sirupsen/logrus"
)

//TODO: NEED REFACTORING. Move to separated module based on logrus_mate!

type LogOutputFormat string

const (
	TextOutputFormat LogOutputFormat = "text"
	JsonOutputFormat LogOutputFormat = "json"
)

type LogConfig struct {
	Format LogOutputFormat
	Level  string
}

const LoggerDefaultConfKey = "logger"

func LoggerGetConfigDefaults() LogConfig {
	defaultLc := LogConfig{
		Format: TextOutputFormat,
		Level:  logrus.InfoLevel.String(),
	}
	return defaultLc
}

func (lc *LogConfig) GetLevelNum() logrus.Level {
	if levelNum, err := logrus.ParseLevel(lc.Level); err == nil {
		return levelNum
	} else {
		return logrus.InfoLevel
	}

}
