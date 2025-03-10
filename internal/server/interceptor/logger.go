package interceptor

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoggerInterceptor struct {
}

func NewLoggerInterceptor() *LoggerInterceptor {
	return &LoggerInterceptor{}
}

func (in *LoggerInterceptor) LoggerInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	st := time.Now()
	log.WithContext(ctx).WithField("request", request).Info("grpc request")
	resp, err := handler(ctx, request)
	in.loggerInterceptorDo(ctx, serverInfo.FullMethod, st, err)
	return resp, err
}

func (*LoggerInterceptor) loggerInterceptorDo(ctx context.Context, fullMethod string, startTime time.Time, err error) {
	st := status.Convert(err)
	var logLevel log.Level
	var logMsg string
	switch st.Code() {
	case codes.OK:
		logLevel = log.InfoLevel
		logMsg = "OK"
	case codes.Unauthenticated, codes.OutOfRange, codes.PermissionDenied, codes.NotFound:
		logLevel = log.InfoLevel
		logMsg = "client error"
	case codes.Internal, codes.Unknown, codes.DataLoss, codes.Unavailable, codes.DeadlineExceeded:
		logLevel = log.ErrorLevel
		logMsg = "server error"
	default:
		logLevel = log.ErrorLevel
		logMsg = "unknown error"
	}
	logger := log.WithContext(ctx).
		WithField("method", fullMethod).
		WithField("elapsed_time", fmt.Sprintf("%.3fms", float64(time.Since(startTime).Microseconds())/1000))
	if err != nil {
		logger.WithField("err", err.Error())
	}
	logger.Log(logLevel, logMsg)
}
