package grpc_util

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ContextKey is the key type of context value.
type ContextKey int

const (
	// The key name used to store username in the context
	// user id is extracted from the jwt token subject field.
	userContextKey ContextKey = iota
	accessTokenContextKey
)

func SetUserContext(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userContextKey, userID)
}

func GetUserID(ctx context.Context) (userId int64, err error) {
	userId, ok := ctx.Value(userContextKey).(int64)
	if !ok {
		err = status.Errorf(codes.Unauthenticated, "unauthenticated")
		return
	}
	return
}

func SetAccessTokenContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, accessTokenContextKey, token)
}

func GetAccessToken(ctx context.Context) (token string, err error) {
	token, ok := ctx.Value(accessTokenContextKey).(string)
	if !ok {
		err = status.Errorf(codes.Unauthenticated, "unauthenticated")
		return
	}
	return
}
