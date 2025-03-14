package interceptor

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
)

// GRPCAuthInterceptor is the auth interceptor for gRPC server.
type GRPCAuthInterceptor struct {
	system system.Service
}

// NewGRPCAuthInterceptor returns a new API auth interceptor.
func NewGRPCAuthInterceptor(i do.Injector) *GRPCAuthInterceptor {
	return &GRPCAuthInterceptor{
		system: do.MustInvoke[system.Service](i),
	}
}

// AuthenticationInterceptor is the unary interceptor for gRPC API.
func (in *GRPCAuthInterceptor) AuthenticationInterceptor(ctx context.Context, request any, serverInfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "failed to parse metadata from incoming context")
	}
	tokenStr, err := getTokenFromMetadata(md)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}
	accessToken, err := in.system.Authenticate(ctx, tokenStr)
	if err != nil {
		if isUnauthorizeAllowedMethod(serverInfo.FullMethod) {
			return handler(ctx, request)
		}
		return nil, err
	}
	user, err := in.system.GetUserByID(ctx, accessToken.UserId)
	if err != nil {
		return nil, err
	}
	if isOnlyForAdminAllowedMethod(serverInfo.FullMethod) && user.Role != model.RoleHost && user.Role != model.RoleAdmin {
		return nil, errors.Errorf("user %d is not admin", user.ID)
	}
	logrus.Infof("%v", user)

	ctx = grpc_util.SetUserContext(ctx, user.ID)
	return handler(ctx, request)
}

func getTokenFromMetadata(md metadata.MD) (string, error) {
	// Check the HTTP request header first.
	authorizationHeaders := md.Get("Authorization")
	if len(md.Get("Authorization")) > 0 {
		authHeaderParts := strings.Fields(authorizationHeaders[0])
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			return "", errors.New("authorization header format must be Bearer {token}")
		}
		return authHeaderParts[1], nil
	}
	// Check the cookie header.
	var accessToken string
	for _, t := range append(md.Get("grpcgateway-cookie"), md.Get("cookie")...) {
		header := http.Header{}
		header.Add("Cookie", t)
		request := http.Request{Header: header}
		if v, _ := request.Cookie("memos.access-token"); v != nil {
			accessToken = v.Value
		}
	}
	return accessToken, nil
}
