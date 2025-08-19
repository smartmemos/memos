package v2

import (
	"context"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

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
func (s *UserService) CreateUser(ctx context.Context, request *connect.Request[v2pb.CreateUserRequest]) (response *connect.Response[modelpb.User], err error) {
	user, err := s.memosService.CreateUser(ctx, &model.CreateUserRequest{
		Username: request.Msg.User.Username,
		Password: request.Msg.User.Password,
	})
	if err != nil {
		return
	}
	response = connect.NewResponse(convertUserToProto(user))
	return
}

func (s *UserService) GetUserStats(ctx context.Context, request *connect.Request[v2pb.GetUserStatsRequest]) (response *connect.Response[v2pb.UserStats], err error) {
	logrus.Info("req: ", request.Msg)
	// userID, err := strconv.ParseInt(req.Msg.Name, 10, 64)
	// if err != nil {
	// 	return
	// }
	// user, err := s.memosService.GetUserByID(ctx, userID)
	// if err != nil {
	// 	return
	// }
	response = connect.NewResponse(&v2pb.UserStats{
		Name: request.Msg.Name,
	})
	return
}

func (s *UserService) GetUserSetting(ctx context.Context, request *connect.Request[v2pb.GetUserSettingRequest]) (response *connect.Response[modelpb.UserSetting], err error) {
	logrus.Info("req: ", request.Msg)
	response = connect.NewResponse(&modelpb.UserSetting{
		Name: request.Msg.Name,
	})
	return
}

func (s *UserService) ListUserSettings(ctx context.Context, request *connect.Request[v2pb.ListUserSettingsRequest]) (response *connect.Response[v2pb.ListUserSettingsResponse], err error) {
	logrus.Info("req: ", request.Msg)

	settings, err := s.memosService.GetUserSettings(ctx, &model.GetUserSettingsRequest{})
	if err != nil {
		return
	}

	response = connect.NewResponse(&v2pb.ListUserSettingsResponse{
		Settings: lo.Map(settings, func(setting *model.UserSetting, _ int) *modelpb.UserSetting {
			return convertUserSettingToProto(setting)
		}),
	})
	return
}

func (s *UserService) ListUserSessions(ctx context.Context, request *connect.Request[v2pb.ListUserSessionsRequest]) (response *connect.Response[v2pb.ListUserSessionsResponse], err error) {
	userID, err := strconv.ParseInt(strings.TrimPrefix(request.Msg.Parent, model.UserNamePrefix), 10, 64)
	if err != nil {
		return
	}
	sessions, err := s.memosService.GetUserSessions(ctx, &model.GetUserSessionsRequest{
		UserID: userID,
	})
	if err != nil {
		return
	}
	response = connect.NewResponse(&v2pb.ListUserSessionsResponse{
		Sessions: lo.Map(sessions, func(session *model.UserSession, _ int) *modelpb.UserSession {
			return convertUserSessionToProto(session)
		}),
	})
	return
}

func (s *UserService) RevokeUserSession(ctx context.Context, request *connect.Request[v2pb.RevokeUserSessionRequest]) (response *connect.Response[emptypb.Empty], err error) {
	parts := strings.Split(request.Msg.Name, "/")
	if len(parts) != 4 || parts[0] != "users" || parts[2] != "sessions" {
		err = errors.Errorf("invalid session name format: %s", request.Msg.Name)
		return
	}
	userID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		err = errors.Errorf("invalid user name: %v", err)
		return
	}
	err = s.memosService.RevokeUserSession(ctx, &model.RevokeUserSessionRequest{
		UserID:    userID,
		SessionID: parts[3],
	})
	if err != nil {
		return
	}
	response = connect.NewResponse(&emptypb.Empty{})
	return
}

func convertUserSessionToProto(session *model.UserSession) *modelpb.UserSession {
	return &modelpb.UserSession{
		Name:             session.Name,
		SessionId:        session.SessionID,
		CreateTime:       timestamppb.New(session.CreateTime),
		LastAccessedTime: timestamppb.New(session.LastAccessedTime),
		ClientInfo: &modelpb.UserSession_ClientInfo{
			UserAgent:  session.ClientInfo.UserAgent,
			IpAddress:  session.ClientInfo.IPAddress,
			DeviceType: session.ClientInfo.DeviceType,
			Os:         session.ClientInfo.OS,
			Browser:    session.ClientInfo.Browser,
		},
	}
}

func convertUserSettingToProto(setting *model.UserSetting) *modelpb.UserSetting {
	return &modelpb.UserSetting{
		Name: setting.Key,
	}
}
