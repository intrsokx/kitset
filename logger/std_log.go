package logger

import (
	"fmt"
	"io"
	"log"
)

var _ Logger = (*stdLogger)(nil)

type stdLogger struct {
	log *log.Logger
}

func NewStdLogger(out io.Writer) Logger {
	return &stdLogger{
		log: log.New(out, "", log.LstdFlags),
	}
}

func (s *stdLogger) Log(level Level, args ...interface{}) {
	buf := defaultPool.Get()
	defer defaultPool.Put(buf)

	buf.WriteString(fmt.Sprintf("[%s] ", level.String()))
	buf.WriteString(fmt.Sprint(args...))

	s.log.Printf(buf.String())
}
