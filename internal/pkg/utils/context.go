package utils

import (
	"context"

	"connectrpc.com/authn"
)

type Info struct {
	UserID    int64
	SessionID string
}

func GetInfo(ctx context.Context) *Info {
	return authn.GetInfo(ctx).(*Info)
}
