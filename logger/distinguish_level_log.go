package logger

import (
	"io"
)

var _ Logger = (*DistinguishLevelLogger)(nil)

//区分日志打印级别logger
type DistinguishLevelLogger struct {
	//大于等于 infoLevel 的log
	infoLog Logger
	//小于 infoLevel 的log
	errLog Logger
}

func (d *DistinguishLevelLogger) Log(level Level, args ...interface{}) {
	if level >= LevelInfo {
		d.infoLog.Log(level, args)
	} else {
		d.errLog.Log(level, args)
	}
}

func NewDistinguishLevelLogger(infoW, errW io.Writer) Logger {
	return &DistinguishLevelLogger{
		infoLog: NewStdLogger(infoW),
		errLog:  NewStdLogger(errW),
	}
}
