package glog

import (
	"context"
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

	return silentLogger
}

type ctxKey int

const (
	ctxLogger ctxKey = iota
)

var silentLogger = new(Silent)

// Silent implements the Logger interface without any output.
type Silent struct{}

func (*Silent) Set(_ ...Field)                      {}
func (*Silent) With(_ ...Field) Logger              { return new(Silent) }
func (*Silent) Trace(_ ...interface{})              {}
func (*Silent) Tracef(_ string, _ ...interface{})   {}
func (*Silent) Debug(_ ...interface{})              {}
func (*Silent) Debugf(_ string, _ ...interface{})   {}
func (*Silent) Info(_ ...interface{})               {}
func (*Silent) Infof(_ string, _ ...interface{})    {}
func (*Silent) Warning(_ ...interface{})            {}
func (*Silent) Warningf(_ string, _ ...interface{}) {}
func (*Silent) Error(_ ...interface{})              {}
func (*Silent) Errorf(_ string, _ ...interface{})   {}
