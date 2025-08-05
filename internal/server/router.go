package server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"

	v2api "github.com/smartmemos/memos/internal/api/v2"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	"github.com/smartmemos/memos/internal/server/interceptor"
	"github.com/smartmemos/memos/internal/server/middleware"
)

func registerHandlers(container do.Injector) http.Handler {
	mux := http.NewServeMux()

	var options []connect.HandlerOption
	options = append(options, connect.WithInterceptors(interceptor.Logger()))
	// options = append(options, connect.WithCompressMinBytes(1024))

	{
		path, authHandler := v2pb.NewAuthServiceHandler(do.MustInvoke[*v2api.AuthService](container), options...)
		mux.Handle(path, authHandler)
	}
	{
		path, authHandler := v2pb.NewUserServiceHandler(do.MustInvoke[*v2api.UserService](container), options...)
		mux.Handle(path, authHandler)
	}
	{
		path, authHandler := v2pb.NewWorkspaceServiceHandler(do.MustInvoke[*v2api.WorkspaceService](container), options...)
		mux.Handle(path, authHandler)
	}
	{
		path, authHandler := v2pb.NewMarkdownServiceHandler(do.MustInvoke[*v2api.MarkdownService](container), options...)
		mux.Handle(path, authHandler)
	}
	{
		path, authHandler := v2pb.NewInboxServiceHandler(do.MustInvoke[*v2api.InboxService](container), options...)
		mux.Handle(path, authHandler)
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
