package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCookie builds a cookie string.
func BuildCookie(ctx context.Context, cookieName, cookieValue, origin string, expireTime time.Time) (string, error) {
	attrs := []string{
		fmt.Sprintf("%s=%s", cookieName, cookieValue),
		"Path=/",
		"HttpOnly",
	}
	if expireTime.IsZero() {
		attrs = append(attrs, "Expires=Thu, 01 Jan 1970 00:00:00 GMT")
	} else {
		attrs = append(attrs, "Expires="+expireTime.Format(time.RFC1123))
	}

	isHTTPS := strings.HasPrefix(origin, "https://")
	if isHTTPS {
		attrs = append(attrs, "SameSite=None")
		attrs = append(attrs, "Secure")
	} else {
		attrs = append(attrs, "SameSite=Strict")
	}
	return strings.Join(attrs, "; "), nil
}

func ClearAuthCookies(ctx context.Context, cookieName, origin string) error {
	sessionCookie, err := BuildCookie(ctx, cookieName, "", origin, time.Time{})
	if err != nil {
		return errors.Wrap(err, "failed to build session cookie")
	}

	// Set both cookies in the response
	if err := grpc.SetHeader(ctx, metadata.New(map[string]string{
		"Set-Cookie": sessionCookie,
	})); err != nil {
		return errors.Wrap(err, "failed to set grpc header")
	}
	return nil
}
