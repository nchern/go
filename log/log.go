package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	DefaultFlags = log.LstdFlags | log.Lshortfile

	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
	levelFatal = "FATAL"
)

var (
	defaultLogger = &defaultImpl{}

	DebugLogger = newLogger(os.Stdout)
	InfoLogger  = newLogger(os.Stdout)
	WarnLogger  = newLogger(os.Stderr)
	ErrorLogger = newLogger(os.Stderr)
)

func newLogger(writer io.Writer) *log.Logger {
	return log.New(writer, "", DefaultFlags)
}

func logIt(l *log.Logger, level, format string, v ...interface{}) {
	msg := ""
	if format != "" {
		msg = fmt.Sprintf(format, v...)
	} else {
		msg = fmt.Sprint(v...)
	}
	l.Output(3, level+" "+msg)
}

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Info(v ...interface{})
	Errorf(format string, v ...interface{})
	Error(v ...interface{})
	Warnf(format string, v ...interface{})
	Warn(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatal(v ...interface{})
}

func L() Logger { return defaultLogger }

type defaultImpl struct{}

func (l *defaultImpl) Debug(v ...interface{}) {
	logIt(DebugLogger, levelDebug, "", v...)
}

func (l *defaultImpl) Debugf(format string, v ...interface{}) {
	logIt(DebugLogger, levelDebug, format, v...)
}

func (l *defaultImpl) Infof(format string, v ...interface{}) {
	logIt(InfoLogger, levelInfo, format, v...)
}

func (l *defaultImpl) Info(v ...interface{}) {
	logIt(InfoLogger, levelInfo, "", v...)
}

func (l *defaultImpl) Errorf(format string, v ...interface{}) {
	logIt(ErrorLogger, levelError, format, v...)
}

func (l *defaultImpl) Error(v ...interface{}) {
	logIt(ErrorLogger, levelError, "", v...)
}

func (l *defaultImpl) Warnf(format string, v ...interface{}) {
	logIt(WarnLogger, levelWarn, format, v...)
}

func (l *defaultImpl) Warn(v ...interface{}) {
	logIt(WarnLogger, levelWarn, "", v...)
}

func (l *defaultImpl) Fatalf(format string, v ...interface{}) {
	logIt(ErrorLogger, levelFatal, format, v...)
	os.Exit(1)
}

func (l *defaultImpl) Fatal(v ...interface{}) {
	logIt(ErrorLogger, levelFatal, "", v...)
	os.Exit(1)
}
