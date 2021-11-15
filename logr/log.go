package logr

import (
	"github.com/newity/glog"
	"github.com/sirupsen/logrus"
)

// Log implements Log interface by using logrus library.
type Log struct {
	l *logrus.Logger
	e *logrus.Entry
}

// New creates a Logger interface implementation.
func New(in *logrus.Logger) *Log {
	return &Log{
		l: in,
		e: logrus.NewEntry(in),
	}
}

// Set updates logger's additional fields.
func (l *Log) Set(fields ...glog.Field) {
	if l.e.Data == nil {
		l.e.Data = make(logrus.Fields)
	}

	for _, field := range fields {
		l.e.Data[field.K] = field.V
	}
}

// With returns a copy of the logger with additional fields.
func (l *Log) With(fields ...glog.Field) glog.Logger {
	f := make(logrus.Fields)

	for k, v := range l.e.Data {
		f[k] = v
	}

	for _, field := range fields {
		f[field.K] = field.V
	}

	return &Log{
		l: l.l,
		e: l.l.WithFields(f),
	}
}

// Trace prints a log message with "trace" log level.
func (l *Log) Trace(args ...interface{}) {
	l.e.Log(logrus.TraceLevel, args...)
}

// Tracef prints a log message with "trace" log level and specified format.
func (l *Log) Tracef(format string, args ...interface{}) {
	l.e.Logf(logrus.TraceLevel, format, args...)
}

// Debug prints a log message with "debug" log level.
func (l *Log) Debug(args ...interface{}) {
	l.e.Log(logrus.DebugLevel, args...)
}

// Debugf prints a log message with "debug" log level and specified format.
func (l *Log) Debugf(format string, args ...interface{}) {
	l.e.Logf(logrus.DebugLevel, format, args...)
}

// Info prints a log message with "info" log level.
func (l *Log) Info(args ...interface{}) {
	l.e.Log(logrus.InfoLevel, args...)
}

// Infof prints a log message with "info" log level and specified format.
func (l *Log) Infof(format string, args ...interface{}) {
	l.e.Logf(logrus.InfoLevel, format, args...)
}

// Warning prints a log message with "warning" log level.
func (l *Log) Warning(args ...interface{}) {
	l.e.Log(logrus.WarnLevel, args...)
}

// Warningf prints a log message with "warning" log level and specified format.
func (l *Log) Warningf(format string, args ...interface{}) {
	l.e.Logf(logrus.WarnLevel, format, args...)
}

// Error prints a log message with "error" log level.
func (l *Log) Error(args ...interface{}) {
	l.e.Log(logrus.ErrorLevel, args...)
}

// Errorf prints a log message with "error" log level and specified format.
func (l *Log) Errorf(format string, args ...interface{}) {
	l.e.Logf(logrus.ErrorLevel, format, args...)
}
