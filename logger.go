package glog

// Field contains key-value parameter for log prefix.
type Field struct {
	K string
	V interface{}
}

// Logger interface is responsible for various logging systems.
type Logger interface {
	// Set updates logger's additional fields.
	Set(fields ...Field)
	// With returns a copy of the logger with additional fields.
	With(fields ...Field) Logger

	// Trace prints a log message with "trace" log level.
	Trace(args ...interface{})
	// Tracef prints a log message with "trace" log level and specified format.
	Tracef(format string, args ...interface{})

	// Debug prints a log message with "debug" log level.
	Debug(args ...interface{})
	// Debugf prints a log message with "debug" log level and specified format.
	Debugf(format string, args ...interface{})

	// Info prints a log message with "info" log level.
	Info(args ...interface{})
	// Infof prints a log message with "info" log level and specified format.
	Infof(format string, args ...interface{})

	// Warning prints a log message with "warning" log level.
	Warning(args ...interface{})
	// Warningf prints a log message with "warning" log level and specified format.
	Warningf(format string, args ...interface{})

	// Error prints a log message with "error" log level.
	Error(args ...interface{})
	// Errorf prints a log message with "error" log level and specified format.
	Errorf(format string, args ...interface{})
}
