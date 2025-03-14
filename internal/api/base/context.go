package base

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
	UserContextKey ContextKey = iota
	AccessTokenContextKey
)

func GetUserID(ctx context.Context) (userId int64, err error) {
	userId, ok := ctx.Value(UserContextKey).(int64)
	if !ok {
		err = status.Errorf(codes.Unauthenticated, "unauthenticated")
		return
	}
	return
}
