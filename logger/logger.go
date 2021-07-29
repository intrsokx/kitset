package logger

type Logger interface {
	Log(level Level, args ...interface{})
}

type logger struct {
	logs []Logger
}

func (l *logger) Log(level Level, args ...interface{}) {
	for _, log := range l.logs {
		log.Log(level, args...)
	}
}

func MultiLogger(logs ...Logger) Logger {
	return &logger{logs: logs}
}
