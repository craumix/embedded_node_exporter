package exporter

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type LogrusTranslator struct {
	LogLevel *logrus.Level
}

func (l *LogrusTranslator) Log(kv ...interface{}) error {
	if len(kv)%2 != 0 {
		return fmt.Errorf("odd number of key-value pairs")
	}

	pairs := make(map[string]any)
	for i := 0; i < len(kv); i += 2 {
		pairs[fmt.Sprintf("%+v", kv[i])] = fmt.Sprintf("%+v", kv[i+1])
	}

	levelMapping := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}

	level := levelMapping[pairs["level"].(string)]
	delete(pairs, "level")

	msg := pairs["msg"]
	delete(pairs, "msg")

	if l.LogLevel == nil || level <= *l.LogLevel {
		logrus.WithFields(pairs).Log(level, msg)
	}

	return nil
}
