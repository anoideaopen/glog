package glog

import (
	"context"
	"os"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

const timeDivider = 1000

// UnaryServerInterceptor returns a new unary server interceptors that adds Logger to the context.
func UnaryServerInterceptor(l Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		var (
			logger = l.With()
			logCtx = NewContext(ctx, logger)
			start  = time.Now()
		)

		addStdFields(ctx, logger, info.FullMethod, start)

		resp, err = handler(
			logCtx,
			req,
		)

		logger.Set(
			Field{K: "grpc.code", V: status.Code(err).String()},
			Field{K: "grpc.time_ms", V: float32(time.Since(start).Nanoseconds()/timeDivider) / timeDivider},
		)

		if err != nil {
			logger.Set(
				Field{K: "error", V: err},
			)
		}

		levelLogf(
			logger,
			status.Code(err),
			"finished unary call with code "+status.Code(err).String())

		return
	}
}

// StreamServerInterceptor returns a new streaming server interceptor that adds Logger to the context.
func StreamServerInterceptor(l Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) (err error) {
		var (
			logger = l.With()
			logCtx = NewContext(stream.Context(), logger)
			start  = time.Now()
		)

		addStdFields(stream.Context(), logger, info.FullMethod, start)

		err = handler(
			srv,
			&streamContextWrapper{stream, logCtx},
		)

		logger.Set(
			Field{K: "grpc.code", V: status.Code(err).String()},
			Field{K: "grpc.time_ms", V: float32(time.Since(start).Nanoseconds()/timeDivider) / timeDivider},
		)

		if err != nil {
			logger.Set(
				Field{K: "error", V: err},
			)
		}

		levelLogf(
			logger,
			status.Code(err),
			"finished streaming call with code "+status.Code(err).String())

		return
	}
}

type streamContextWrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *streamContextWrapper) Context() context.Context {
	return w.ctx
}

// ReplaceGrpcLogger sets the given Logger as a gRPC-level logger v2.
// This should be called *before* any other initialization, preferably from init() functions.
func ReplaceGrpcLogger(l Logger) {
	grpclog.SetLoggerV2(
		&depthLoggerWrapper{
			l: l.With(Field{K: "module", V: "system"}),
		},
	)
}

type depthLoggerWrapper struct {
	l Logger
}

// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
func (w *depthLoggerWrapper) Info(args ...interface{}) {
	w.l.Info(args...)
}

// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
func (w *depthLoggerWrapper) Infoln(args ...interface{}) {
	w.Info(args...)
}

// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
func (w *depthLoggerWrapper) Infof(format string, args ...interface{}) {
	w.l.Infof(format, args...)
}

// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
func (w *depthLoggerWrapper) Warning(args ...interface{}) {
	w.l.Warning(args...)
}

// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
func (w *depthLoggerWrapper) Warningln(args ...interface{}) {
	w.Warning(args...)
}

// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
func (w *depthLoggerWrapper) Warningf(format string, args ...interface{}) {
	w.l.Warningf(format, args...)
}

// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
func (w *depthLoggerWrapper) Error(args ...interface{}) {
	w.l.Error(args...)
}

// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
func (w *depthLoggerWrapper) Errorln(args ...interface{}) {
	w.Error(args...)
}

// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
func (w *depthLoggerWrapper) Errorf(format string, args ...interface{}) {
	w.l.Errorf(format, args...)
}

// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (w *depthLoggerWrapper) Fatal(args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			w.l.Error(r)
			os.Exit(1)
		}
	}()
	w.l.Error(args...)
	panic(nil)
}

// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (w *depthLoggerWrapper) Fatalln(args ...interface{}) {
	w.Fatal(args...)
}

// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
// gRPC ensures that all Fatal logs will exit with os.Exit(1).
// Implementations may also call os.Exit() with a non-zero exit code.
func (w *depthLoggerWrapper) Fatalf(format string, args ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			w.l.Error(r)
			os.Exit(1)
		}
	}()
	w.l.Errorf(format, args...)
	panic(nil)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (w *depthLoggerWrapper) V(int) bool {
	return true
}

func addStdFields(ctx context.Context, logger Logger, fullMethodString string, start time.Time) {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)

	logger.Set(
		Field{K: "system", V: "grpc"},
		Field{K: "span.kind", V: "server"},
		Field{K: "grpc.service", V: service},
		Field{K: "grpc.method", V: method},
		Field{K: "grpc.start_time", V: start.Format(time.RFC3339Nano)},
	)

	if d, ok := ctx.Deadline(); ok {
		logger.Set(
			Field{K: "grpc.request.deadline", V: d.Format(time.RFC3339)},
		)
	}
}

func levelLogf(logger Logger, code codes.Code, format string, args ...interface{}) {
	switch code {
	case
		codes.OK,
		codes.Canceled,
		codes.InvalidArgument,
		codes.NotFound,
		codes.AlreadyExists,
		codes.Unauthenticated:
		logger.Infof(format, args...)
	case
		codes.DeadlineExceeded,
		codes.PermissionDenied,
		codes.ResourceExhausted,
		codes.FailedPrecondition,
		codes.Aborted,
		codes.OutOfRange,
		codes.Unavailable:
		logger.Warningf(format, args...)
	case
		codes.Unknown,
		codes.Unimplemented,
		codes.Internal,
		codes.DataLoss:
		logger.Errorf(format, args...)
	default:
		logger.Errorf(format, args...)
	}
}
