package grpc_util

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// AccessTokenCookieName is the cookie name of access token.
	AccessTokenCookieName = "memos.access-token"
)

func ClearAccessTokenCookie(ctx context.Context) error {
	return SetAccessTokenCookie(ctx, "", time.Time{})
}

func SetAccessTokenCookie(ctx context.Context, token string, expireTime time.Time) (err error) {
	cookie, err := buildAccessTokenCookie(ctx, token, expireTime)
	if err != nil {
		err = errors.Errorf("failed to build access token cookie, err: %s", err)
		return
	}
	if err = grpc.SetHeader(ctx, metadata.New(map[string]string{"Set-Cookie": cookie})); err != nil {
		err = errors.Errorf("failed to set grpc header, error: %v", err)
		return
	}
	return
}

func buildAccessTokenCookie(ctx context.Context, accessToken string, expireTime time.Time) (string, error) {
	attrs := []string{
		fmt.Sprintf("%s=%s", AccessTokenCookieName, accessToken),
		"Path=/",
		"HttpOnly",
	}
	if expireTime.IsZero() {
		attrs = append(attrs, "Expires=Thu, 01 Jan 1970 00:00:00 GMT")
	} else {
		attrs = append(attrs, "Expires="+expireTime.Format(time.RFC1123))
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("failed to get metadata from context")
	}
	var origin string
	for _, v := range md.Get("origin") {
		origin = v
	}
	if strings.HasPrefix(origin, "https://") {
		attrs = append(attrs, "SameSite=None")
		attrs = append(attrs, "Secure")
	} else {
		attrs = append(attrs, "SameSite=Strict")
	}
	return strings.Join(attrs, "; "), nil
}
