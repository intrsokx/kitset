package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestMultiLogger(t *testing.T) {
	f1, _ := os.Create("log1")
	f2, _ := os.Create("log2")
	defer f1.Close()
	defer f2.Close()

	log1 := NewStdLogger(f1)
	log2 := NewStdLogger(f2)

	log := MultiLogger(log1, log2)
	log.Log(LevelInfo, "this is info")
	log.Log(LevelWarn, "this is warn")
}

func TestNewDistinguishLevelLogger(t *testing.T) {
	f1, _ := os.Create("info.log")
	f2, _ := os.Create("error.log")
	defer f1.Close()
	defer f2.Close()

	log := NewDistinguishLevelLogger(f1, f2)

	log.Log(LevelInfo, "info msg")
	log.Log(LevelWarn, "warn msg")
	log.Log(LevelError, "error msg")
	log.Log(LevelDebug, "debug msg")
	log.Log(LevelFatal, "fatal msg")
}

func TestNewLogrusLogger(t *testing.T) {
	lgrs := logrus.New()
	lgrs.SetOutput(os.Stdout)
	lgrs.SetReportCaller(true)
	//lgrs.SetLevel(logrus.DebugLevel)

	log := NewLogrusLogger(lgrs)
	log.Log(LevelInfo, "info msg")
	//need set lgrs level
	log.Log(LevelDebug, "debug msg")
	log.Log(LevelError, "error msg")
}
