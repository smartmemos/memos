package v2

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"github.com/sirupsen/logrus"

	"github.com/smartmemos/memos/internal/memos"
	"github.com/smartmemos/memos/internal/memos/model"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type UserService struct {
	v2pb.UnimplementedUserServiceHandler
	memosService memos.Service
}

func NewUserService(i do.Injector) (*UserService, error) {
	return &UserService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, req *connect.Request[v2pb.CreateUserRequest]) (resp *connect.Response[modelpb.User], err error) {
	user, err := s.memosService.CreateUser(ctx, &model.CreateUserRequest{
		Username: req.Msg.User.Username,
		Password: req.Msg.User.Password,
	})
	if err != nil {
		return
	}
	resp = connect.NewResponse(convertUserToProto(user))
	return
}

func (s *UserService) GetUserStats(ctx context.Context, req *connect.Request[v2pb.GetUserStatsRequest]) (resp *connect.Response[v2pb.UserStats], err error) {
	logrus.Info("req: ", req.Msg)
	// userID, err := strconv.ParseInt(req.Msg.Name, 10, 64)
	// if err != nil {
	// 	return
	// }
	// user, err := s.memosService.GetUserByID(ctx, userID)
	// if err != nil {
	// 	return
	// }
	resp = connect.NewResponse(&v2pb.UserStats{
		Name: req.Msg.Name,
	})
	return
}
