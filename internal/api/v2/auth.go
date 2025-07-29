package v2

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type AuthService struct {
	v2pb.UnimplementedAuthServiceHandler
	memosService memos.Service
}

func NewAuthService(i do.Injector) (*AuthService, error) {
	return &AuthService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

// SignUp creates a new user.
func (s *AuthService) SignUp(ctx context.Context, req *connect.Request[v2pb.SignUpRequest]) (resp *connect.Response[modelpb.User], err error) {
	user, err := s.memosService.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Msg.Username,
		Password: req.Msg.Password,
	})
	if err != nil {
		return
	}
	resp = connect.NewResponse(convertUserToProto(user))
	return
}

func convertUserToProto(user *model.User) *modelpb.User {
	return &modelpb.User{
		Id:          user.ID,
		Name:        fmt.Sprintf("users/%d", user.ID),
		Username:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		AvatarUrl:   user.AvatarURL,
		Description: user.Description,
		CreateAt:    timestamppb.New(user.CreatedAt),
		UpdateAt:    timestamppb.New(user.UpdatedAt),
		Role:        modelpb.User_Role(modelpb.User_Role_value[string(user.Role)]),
		State:       modelpb.State(modelpb.State_value[string(user.Status)]),
	}
}
