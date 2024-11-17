package utils

import (
	"io"
	"log"
)

type LogLevel int

const (
	LevelError LogLevel = iota
	LevelWarn
	LevelDebug
)

type Logger interface {
	SetLevel(level LogLevel)
	Errorf(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
}

func NewLogger(output io.Writer, prefix string, flag int) Logger {
	return &logger{l: log.New(output, prefix, flag), level: LevelDebug}
}

type logger struct {
	l     *log.Logger
	level LogLevel
}

func (l *logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if l.level >= LevelError {
		l.output("ERROR", format, v...)
	}
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if l.level >= LevelWarn {
		l.output("WARN", format, v...)
	}
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if l.level >= LevelDebug {
		l.output("DEBUG", format, v...)
	}
}

func (l *logger) output(level, format string, v ...interface{}) {
	format = level + " -> " + format
	if len(v) == 0 {
		l.l.Print(format)
		return
	}
	l.l.Printf(format, v...)
}
