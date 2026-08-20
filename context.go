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

func (*Silent) Set(_ ...Field)              {}
func (*Silent) With(_ ...Field) Logger      { return new(Silent) }
func (*Silent) Trace(_ ...any)              {}
func (*Silent) Tracef(_ string, _ ...any)   {}
func (*Silent) Debug(_ ...any)              {}
func (*Silent) Debugf(_ string, _ ...any)   {}
func (*Silent) Info(_ ...any)               {}
func (*Silent) Infof(_ string, _ ...any)    {}
func (*Silent) Warning(_ ...any)            {}
func (*Silent) Warningf(_ string, _ ...any) {}
func (*Silent) Error(_ ...any)              {}
func (*Silent) Errorf(_ string, _ ...any)   {}
