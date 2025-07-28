package server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	apiv1 "github.com/smartmemos/memos/internal/api/v1"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	"github.com/smartmemos/memos/internal/server/interceptor"
	"github.com/smartmemos/memos/internal/server/middleware"
)

func registerHandlers(container do.Injector) http.Handler {
	mux := http.NewServeMux()

	var options []connect.HandlerOption
	options = append(options, connect.WithInterceptors(interceptor.Logger()))
	// options = append(options, connect.WithCompressMinBytes(1024))

	{
		path, authHandler := v1pb.NewAuthServiceHandler(do.MustInvoke[*apiv1.AuthService](container), options...)
		mux.Handle(path, authHandler)
	}
	{
		path, memoHandler := v1pb.NewMemoServiceHandler(do.MustInvoke[*apiv1.MemoService](container), options...)
		mux.Handle(path, wrapHandler(memoHandler))
	}
	{
		path, userHandler := v1pb.NewUserServiceHandler(do.MustInvoke[*apiv1.UserService](container), options...)
		mux.Handle(path, wrapHandler(userHandler))
	}
	{
		path, workspaceHandler := v1pb.NewWorkspaceServiceHandler(do.MustInvoke[*apiv1.WorkspaceService](container), options...)
		mux.Handle(path, wrapHandler(workspaceHandler))
	}

	handler := wrapHandler(mux, middleware.CORS, middleware.NewAuth(container).Auth)
	return handler
}

func wrapHandler(handler http.Handler, chains ...func(http.Handler) http.Handler) http.Handler {
	// 顺序很重要，最外层的最先执行
	for i := len(chains) - 1; i >= 0; i-- {
		handler = chains[i](handler)
	}
	return handler
}
