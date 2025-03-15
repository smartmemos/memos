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

	v1api "github.com/smartmemos/memos/internal/api/v1"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
)

func (s *Server) registerGateway(ctx context.Context, container do.Injector) error {
	v1pb.RegisterAuthServiceServer(s.grpcServer, do.MustInvoke[*v1api.AuthService](container))
	v1pb.RegisterMemoServiceServer(s.grpcServer, do.MustInvoke[*v1api.MemoService](container))
	v1pb.RegisterUserServiceServer(s.grpcServer, do.MustInvoke[*v1api.UserService](container))
	v1pb.RegisterWorkspaceServiceServer(s.grpcServer, do.MustInvoke[*v1api.WorkspaceService](container))

	conn, err := grpc.NewClient(s.profile.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)),
	)
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux()
	if err = v1pb.RegisterAuthServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = v1pb.RegisterMemoServiceHandler(ctx, mux, conn); err != nil {
		return err
	}
	if err = v1pb.RegisterUserServiceHandler(ctx, mux, conn); err != nil {
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
