package v1

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/api/base"
	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/model"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	systempb "github.com/smartmemos/memos/internal/proto/model/system"
)

type AuthService struct {
	v1pb.UnimplementedAuthServiceServer
	system system.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		system: do.MustInvoke[system.Service](i),
	}, nil
}

func (s *AuthService) SignIn(ctx context.Context, req *v1pb.SignInRequest) (resp *systempb.User, err error) {
	accessToken, err := s.system.SignIn(ctx, &model.SignInRequest{
		Username:    req.Username,
		Password:    req.Password,
		NeverExpire: req.NeverExpire,
	})
	if err != nil {
		return
	}
	if err = s.setAccessTokenCookie(ctx, accessToken.Token, accessToken.ExpiresAt); err != nil {
		return
	}
	user, err := s.system.GetUserByID(ctx, accessToken.UserId)
	if err != nil {
		return
	}
	resp = convertUserToProto(user)
	return resp, nil
}

func (s *AuthService) GetAuthStatus(ctx context.Context, req *v1pb.GetAuthStatusRequest) (resp *systempb.User, err error) {
	userID, err := base.GetUserID(ctx)
	if err != nil {
		return
	}
	user, err := s.system.GetUserByID(ctx, userID)
	if err != nil {
		if err = s.clearAccessTokenCookie(ctx); err != nil {
			err = status.Errorf(codes.Internal, "failed to set grpc header: %v", err)
		} else {
			err = status.Errorf(codes.Unauthenticated, "user not found")
		}
		return
	}
	resp = convertUserToProto(user)
	return
}

func (s *AuthService) clearAccessTokenCookie(ctx context.Context) error {
	return s.setAccessTokenCookie(ctx, "", time.Time{})
}

func (s *AuthService) setAccessTokenCookie(ctx context.Context, token string, expireTime time.Time) (err error) {
	cookie, err := s.buildAccessTokenCookie(ctx, token, expireTime)
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

func (s *AuthService) buildAccessTokenCookie(ctx context.Context, accessToken string, expireTime time.Time) (string, error) {
	attrs := []string{
		fmt.Sprintf("%s=%s", base.AccessTokenCookieName, accessToken),
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
