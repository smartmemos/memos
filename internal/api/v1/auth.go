package v1

import (
	"context"

	"github.com/samber/do/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartmemos/memos/internal/module/system"
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

func (s *AuthService) SignIn(context.Context, *v1pb.SignInRequest) (*systempb.User, error) {
	// s.system.SignIn()
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
