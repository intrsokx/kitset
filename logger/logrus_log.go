package logger

import "github.com/sirupsen/logrus"

var _ Logger = (*LogrusLogger)(nil)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger(log *logrus.Logger) Logger {
	return &LogrusLogger{log: log}
}

func (l *LogrusLogger) Log(level Level, args ...interface{}) {
	switch level {
	case LevelDebug:
		l.log.Debug(args...)
	case LevelInfo:
		l.log.Info(args...)
	case LevelWarn:
		l.log.Warn(args...)
	case LevelError:
		l.log.Error(args...)
	case LevelFatal:
		l.log.Fatal(args...)
	default:
		l.log.Info(args...)
	}
}
