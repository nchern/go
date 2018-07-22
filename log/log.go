package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	// DefaultFlags to use for log package
	DefaultFlags = log.LstdFlags | log.Lshortfile

	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
	levelFatal = "FATAL"
)

var (
	defaultLogger = &defaultImpl{}

	// DebugLogSink accepts debug entries
	DebugLogSink = newLogSink(os.Stdout)

	// InfoLogSink accepts info entries
	InfoLogSink = newLogSink(os.Stdout)

	// WarnLogSink accepts warn entries
	WarnLogSink = newLogSink(os.Stderr)

	// ErrorLogSink accepts error entries
	ErrorLogSink = newLogSink(os.Stderr)
)

func newLogSink(writer io.Writer) *log.Logger {
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

// Logger defines the logger interface
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

// L is a shortcut to return default logger
func L() Logger { return defaultLogger }

type defaultImpl struct{}

func (l *defaultImpl) Debug(v ...interface{}) {
	logIt(DebugLogSink, levelDebug, "", v...)
}

func (l *defaultImpl) Debugf(format string, v ...interface{}) {
	logIt(DebugLogSink, levelDebug, format, v...)
}

func (l *defaultImpl) Infof(format string, v ...interface{}) {
	logIt(InfoLogSink, levelInfo, format, v...)
}

func (l *defaultImpl) Info(v ...interface{}) {
	logIt(InfoLogSink, levelInfo, "", v...)
}

func (l *defaultImpl) Errorf(format string, v ...interface{}) {
	logIt(ErrorLogSink, levelError, format, v...)
}

func (l *defaultImpl) Error(v ...interface{}) {
	logIt(ErrorLogSink, levelError, "", v...)
}

func (l *defaultImpl) Warnf(format string, v ...interface{}) {
	logIt(WarnLogSink, levelWarn, format, v...)
}

func (l *defaultImpl) Warn(v ...interface{}) {
	logIt(WarnLogSink, levelWarn, "", v...)
}

func (l *defaultImpl) Fatalf(format string, v ...interface{}) {
	logIt(ErrorLogSink, levelFatal, format, v...)
	os.Exit(1)
}

func (l *defaultImpl) Fatal(v ...interface{}) {
	logIt(ErrorLogSink, levelFatal, "", v...)
	os.Exit(1)
}
