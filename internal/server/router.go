package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/samber/do/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	memov1 "github.com/smartmemos/memos/internal/api/v1/memo"
	v1 "github.com/smartmemos/memos/internal/proto/api/v1"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
)

func (s *Server) registerGateway(ctx context.Context, container do.Injector) error {
	v1pb.RegisterMemoServiceServer(s.grpcServer, do.MustInvoke[*memov1.Service](container))
	v1.RegisterMemoServiceServer(s.grpcServer, do.MustInvoke[*memov1.Service](container))

	conn, err := grpc.NewClient(s.profile.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)),
	)
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux()
	if err = v1pb.RegisterMemoServiceHandler(ctx, mux, conn); err != nil {
		return err
	}

	s.echoServer.Any("/api/*", echo.WrapHandler(mux), middleware.CORS())

	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(_ string) bool {
			return true
		}),
	}
	wrappedGrpc := grpcweb.WrapServer(s.grpcServer, options...)
	s.echoServer.Any("/api.v1.*", echo.WrapHandler(wrappedGrpc))
	return nil
}
