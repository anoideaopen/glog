package std

import (
	"fmt"
	"log"
	"strings"

	"github.com/newity/glog"
)

// Log implements the Logger interface by using standard golang logger.
type Log struct {
	lvl    Level
	l      *log.Logger
	fields map[string]interface{}
	fdata  string
}

// New creates a Logger interface implementation.
func New(in *log.Logger, lvl Level) *Log {
	return &Log{
		lvl: lvl,
		l:   in,
	}
}

// SetLevel updates log level.
func (l *Log) SetLevel(lvl Level) *Log {
	l.lvl = lvl

	return l
}

// Set updates logger's additional fields.
func (l *Log) Set(fields ...glog.Field) {
	l.updateFields(fields...)
}

// With returns a copy of the logger with additional fields.
func (l *Log) With(fields ...glog.Field) glog.Logger {
	lcopy := New(l.l, l.lvl)
	lcopy.fields = make(map[string]interface{})

	for k, v := range l.fields {
		lcopy.fields[k] = v
	}

	return lcopy.updateFields(fields...)
}

// Print prints log message with a specified level.
func (l *Log) Print(lvl Level, args ...interface{}) {
	if l.lvl < lvl {
		return
	}

	l.l.Print(append([]interface{}{"[" + lvl.String() + "]" + l.fdata}, args...)...)
}

// Printf prints log message with a specified level and format.
func (l *Log) Printf(lvl Level, format string, args ...interface{}) {
	if l.lvl < lvl {
		return
	}

	l.l.Printf("["+lvl.String()+"]"+l.fdata+format, args...)
}

// Trace prints a log message with "trace" log level.
func (l *Log) Trace(args ...interface{}) {
	l.Print(LevelTrace, args...)
}

// Tracef prints a log message with "trace" log level and specified format.
func (l *Log) Tracef(format string, args ...interface{}) {
	l.Printf(LevelTrace, format, args...)
}

// Debug prints a log message with "debug" log level.
func (l *Log) Debug(args ...interface{}) {
	l.Print(LevelDebug, args...)
}

// Debugf prints a log message with "debug" log level and specified format.
func (l *Log) Debugf(format string, args ...interface{}) {
	l.Printf(LevelDebug, format, args...)
}

// Info prints a log message with "info" log level.
func (l *Log) Info(args ...interface{}) {
	l.Print(LevelInfo, args...)
}

// Infof prints a log message with "info" log level and specified format.
func (l *Log) Infof(format string, args ...interface{}) {
	l.Printf(LevelInfo, format, args...)
}

// Warning prints a log message with "warning" log level.
func (l *Log) Warning(args ...interface{}) {
	l.Print(LevelWarning, args...)
}

// Warningf prints a log message with "warning" log level and specified format.
func (l *Log) Warningf(format string, args ...interface{}) {
	l.Printf(LevelWarning, format, args...)
}

// Error prints a log message with "error" log level.
func (l *Log) Error(args ...interface{}) {
	l.Print(LevelError, args...)
}

// Errorf prints a log message with "error" log level and specified format.
func (l *Log) Errorf(format string, args ...interface{}) {
	l.Printf(LevelError, format, args...)
}

func (l *Log) updateFields(fields ...glog.Field) glog.Logger {
	if l.fields == nil {
		l.fields = make(map[string]interface{})
	}

	for _, field := range fields {
		l.fields[field.K] = field.V
	}

	sb := strings.Builder{}
	for k, v := range l.fields {
		sb.WriteString(fmt.Sprintf("%s:%v ", k, v))
	}

	out := sb.String()
	l.fdata = " {" + out[:len(out)-1] + "} "

	return l
}
