package auth

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	"github.com/smartmemos/memos/internal/proto/model/system"
)

func (s *Service) SignIn(context.Context, *v1pb.SignInRequest) (*system.User, error) {
	// s.system.SignIn()
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
