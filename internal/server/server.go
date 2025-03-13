package server

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/labstack/echo"
	"github.com/samber/do/v2"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"

	"github.com/smartmemos/memos/internal/server/interceptor"
)

type Server struct {
	profile    *Profile
	grpcServer *grpc.Server
	echoServer *echo.Echo
}

func NewServer(profile *Profile) (*Server, error) {
	s := &Server{
		profile: profile,
		grpcServer: grpc.NewServer(
			grpc.MaxRecvMsgSize(100*1024*1024),
			grpc.ChainUnaryInterceptor(
				grpc_recovery.UnaryServerInterceptor(),
				interceptor.NewLoggerInterceptor().LoggerInterceptor,
				interceptor.NewGRPCAuthInterceptor(cfg.Key, nil).AuthenticationInterceptor,
			),
		),
	}
	echoServer := echo.New()
	echoServer.Debug = true
	echoServer.HideBanner = true
	s.echoServer = echoServer
	return s, nil
}

func (s *Server) Start(ctx context.Context, container do.Injector) error {
	listener, err := net.Listen("tcp", s.profile.Addr)
	if err != nil {
		return err
	}
	if err = s.registerGateway(ctx, container); err != nil {
		return err
	}
	m := cmux.New(listener)
	// 启动gRPC服务
	go func() {
		grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
		log.Info("starting gRPC server")
		if err = s.grpcServer.Serve(grpcListener); err != nil && err != cmux.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	// 启动http服务
	go func() {
		s.echoServer.Listener = m.Match(cmux.HTTP1Fast())
		log.Info("starting HTTP server")
		if err = s.echoServer.Start(s.profile.Addr); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	go func() {
		if err = m.Serve(); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			log.Errorf("failed to start server: %v", err)
		}
	}()
	return nil
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Info("shutdown...")
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	if err := s.echoServer.Shutdown(ctx); err != nil {
		log.Errorf("failed to shutdown server, error: %v\n", err)
	}
	s.grpcServer.GracefulStop()
}
