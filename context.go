package glog

import (
	"context"
	"fmt"
)

// NewContext adds Logger to the Context.
func NewContext(parent context.Context, l Logger) context.Context {
	return context.WithValue(parent, ctxLogger, l)
}

// FromContext gets Logger from the Context.
func FromContext(ctx context.Context) Logger {
	if val, ok := ctx.Value(ctxLogger).(Logger); ok {
		return val
	}

	return mockLogger
}

type ctxKey int

const (
	ctxLogger ctxKey = iota
)

var mockLogger = new(emptyLogger)

type emptyLogger struct{}

func (l *emptyLogger) Set(_ ...Field)                            {}
func (l *emptyLogger) With(_ ...Field) Logger                    { return l }
func (l *emptyLogger) Trace(_ ...interface{})                    {}
func (l *emptyLogger) Tracef(_ string, _ ...interface{})         {}
func (l *emptyLogger) Debug(_ ...interface{})                    {}
func (l *emptyLogger) Debugf(_ string, _ ...interface{})         {}
func (l *emptyLogger) Info(_ ...interface{})                     {}
func (l *emptyLogger) Infof(_ string, _ ...interface{})          {}
func (l *emptyLogger) Warning(_ ...interface{})                  {}
func (l *emptyLogger) Warningf(_ string, _ ...interface{})       {}
func (l *emptyLogger) Error(_ ...interface{})                    {}
func (l *emptyLogger) Errorf(_ string, _ ...interface{})         {}
func (l *emptyLogger) Panic(args ...interface{})                 { panic(fmt.Sprint(args...)) }
func (l *emptyLogger) Panicf(format string, args ...interface{}) { panic(fmt.Sprintf(format, args...)) }
