package glog

// Field contains key-value parameter for log prefix.
type Field struct {
	K string
	V any
}

// Logger interface is responsible for various logging systems.
type Logger interface { //nolint:interfacebloat
	// Set updates logger's additional fields.
	Set(fields ...Field)
	// With returns a copy of the logger with additional fields.
	With(fields ...Field) Logger

	// Trace prints a log message with "trace" log level.
	Trace(args ...any)
	// Tracef prints a log message with "trace" log level and specified format.
	Tracef(format string, args ...any)

	// Debug prints a log message with "debug" log level.
	Debug(args ...any)
	// Debugf prints a log message with "debug" log level and specified format.
	Debugf(format string, args ...any)

	// Info prints a log message with "info" log level.
	Info(args ...any)
	// Infof prints a log message with "info" log level and specified format.
	Infof(format string, args ...any)

	// Warning prints a log message with "warning" log level.
	Warning(args ...any)
	// Warningf prints a log message with "warning" log level and specified format.
	Warningf(format string, args ...any)

	// Error prints a log message with "error" log level.
	Error(args ...any)
	// Errorf prints a log message with "error" log level and specified format.
	Errorf(format string, args ...any)
}
