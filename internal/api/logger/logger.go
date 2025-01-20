package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	ERROR = "ERROR"
	INFO  = "INFO"
)

var (
	instance *logger
	once     sync.Once
)

type logger struct{}

type Logger interface {
	Info(msg string)
	Error(msg string)
}

func GetInstance() Logger {
	once.Do(func() {
		instance = &logger{}
	})
	return instance
}

func log(level string, msg string) {
	now := time.Now()
	dateTime := fmt.Sprintf("%s.%03d", now.Format("2006-01-02 15:04:05"), now.Nanosecond()/1e6)
	logMsg := fmt.Sprintf("%s\t[%s] %s\n", level, dateTime, msg)
	switch level {
	case INFO:
		fmt.Fprintln(os.Stdout, logMsg)
	case ERROR:
		fmt.Fprintln(os.Stderr, logMsg)
	}
}

func (l *logger) Info(msg string) {
	log(INFO, msg)
}

func (l *logger) Error(msg string) {
	log(ERROR, msg)
}
