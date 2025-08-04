package v2

import (
	"context"

	"connectrpc.com/connect"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
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

func (s *UserService) GetUserSetting(ctx context.Context, req *connect.Request[v2pb.GetUserSettingRequest]) (resp *connect.Response[modelpb.UserSetting], err error) {
	logrus.Info("req: ", req.Msg)
	resp = connect.NewResponse(&modelpb.UserSetting{
		Name: req.Msg.Name,
	})
	return
}

func (s *UserService) ListUserSettings(ctx context.Context, req *connect.Request[v2pb.ListUserSettingsRequest]) (resp *connect.Response[v2pb.ListUserSettingsResponse], err error) {
	logrus.Info("req: ", req.Msg)

	settings, err := s.memosService.GetUserSettings(ctx, &model.GetUserSettingsRequest{})
	if err != nil {
		return
	}

	resp = connect.NewResponse(&v2pb.ListUserSettingsResponse{
		Settings: lo.Map(settings, func(setting *model.UserSetting, _ int) *modelpb.UserSetting {
			return convertUserSettingToProto(setting)
		}),
	})
	return
}

func convertUserSettingToProto(setting *model.UserSetting) *modelpb.UserSetting {
	return &modelpb.UserSetting{
		Name: setting.Key,
	}
}
