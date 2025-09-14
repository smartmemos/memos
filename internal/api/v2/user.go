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
	"github.com/smartmemos/memos/internal/pkg/utils"
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

// UpdateUser updates a user.
func (s *UserService) UpdateUser(ctx context.Context, request *connect.Request[v2pb.UpdateUserRequest]) (response *connect.Response[modelpb.User], err error) {
	logrus.Info("req: ", request.Msg)

	userID, err := strconv.ParseInt(strings.TrimPrefix(request.Msg.User.Name, model.UserNamePrefix), 10, 64)
	if err != nil {
		return
	}
	userInfo := utils.GetInfo(ctx)
	if userInfo == nil {
		err = errors.New("failed to get user")
		return
	} else if userID != userInfo.UserID {
		err = errors.New("user not found")
		return
	}

	user, err := s.memosService.UpdateUser(ctx, &model.UpdateUserRequest{
		UpdateMask:  request.Msg.UpdateMask.Paths,
		ID:          userID,
		Username:    request.Msg.User.Username,
		Role:        model.Role(request.Msg.User.Role),
		Email:       request.Msg.User.Email,
		Nickname:    request.Msg.User.Nickname,
		Password:    request.Msg.User.Password,
		AvatarURL:   request.Msg.User.AvatarUrl,
		Description: request.Msg.User.Description,
		// Status:      model.RowStatus(request.Msg.User.Status),
	})
	if err != nil {
		return
	}
	response = connect.NewResponse(convertUserToProto(user))
	return
}

func (s *UserService) GetUser(ctx context.Context, request *connect.Request[v2pb.GetUserRequest]) (response *connect.Response[modelpb.User], err error) {
	logrus.Info("req: ", request.Msg)
	userID, err := strconv.ParseInt(strings.TrimPrefix(request.Msg.Name, model.UserNamePrefix), 10, 64)
	if err != nil {
		return
	}
	user, err := s.memosService.GetUserByID(ctx, userID)
	if err != nil {
		return
	}
	response = connect.NewResponse(convertUserToProto(user))
	return
}

func (s *UserService) SearchUsers(ctx context.Context, request *connect.Request[v2pb.SearchUsersRequest]) (response *connect.Response[v2pb.SearchUsersResponse], err error) {
	// logrus.Info("req: ", request.Msg)
	// users, err := s.memosService.SearchUsers(ctx, &model.SearchUsersRequest{
	// 	Query:     request.Msg.Query,
	// 	PageSize:  request.Msg.PageSize,
	// 	PageToken: request.Msg.PageToken,
	// })
	return
}

func (s *UserService) ListUsers(ctx context.Context, request *connect.Request[v2pb.ListUsersRequest]) (response *connect.Response[v2pb.ListUsersResponse], err error) {
	logrus.Info("req: ", request.Msg)
	// users, err := s.memosService.ListUsers(ctx, &model.ListUsersRequest{
	// 	PageSize:  request.Msg.PageSize,
	// 	PageToken: request.Msg.PageToken,
	// 	Filter:    request.Msg.Filter,
	// 	OrderBy:   request.Msg.OrderBy,
	// 	ShowDeleted: request.Msg.ShowDeleted,
	// })
	return
}

// GetUserStats returns statistics for a specific user.
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

func (s *UserService) UpdateUserSetting(ctx context.Context, request *connect.Request[v2pb.UpdateUserSettingRequest]) (response *connect.Response[modelpb.UserSetting], err error) {
	logrus.Info("req: ", request.Msg)

	parts := strings.Split(request.Msg.Setting.Name, "/")
	if len(parts) != 4 || parts[0] != "users" || parts[2] != "settings" {
		err = errors.Errorf("invalid setting name format: %s", request.Msg.Setting.Name)
		return
	}
	userID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		err = errors.Errorf("invalid user name: %v", err)
		return
	}
	userInfo := utils.GetInfo(ctx)
	if userInfo == nil {
		err = errors.New("failed to get user")
		return
	} else if userID != userInfo.UserID {
		err = errors.New("user not found")
		return
	}
	var settingValue model.UserSettingValue
	if request.Msg.Setting.Value != nil {
		switch value := request.Msg.Setting.Value.(type) {
		case *modelpb.UserSetting_GeneralSetting_:
			settingValue = model.UserSettingValue{
				GeneralUserSetting: &model.GeneralUserSetting{
					Locale:         value.GeneralSetting.Locale,
					Appearance:     value.GeneralSetting.Appearance,
					MemoVisibility: value.GeneralSetting.MemoVisibility,
					Theme:          value.GeneralSetting.Theme,
				},
			}
		}
	}
	setting, err := s.memosService.UpdateUserSetting(ctx, &model.UpdateUserSettingRequest{
		UpdateMask: request.Msg.UpdateMask.Paths,
		UserID:     userID,
		Key:        model.UserSettingKey(parts[3]),
		Value:      settingValue,
	})
	if err != nil {
		return
	}

	response = connect.NewResponse(convertUserSettingToProto(setting))
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

func (s *UserService) ListUserAccessTokens(ctx context.Context, request *connect.Request[v2pb.ListUserAccessTokensRequest]) (response *connect.Response[v2pb.ListUserAccessTokensResponse], err error) {
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

func convertUserAccessToken(token *model.UserAccessToken) *modelpb.UserAccessToken {
	return &modelpb.UserAccessToken{
		Name:        token.Name,
		AccessToken: token.AccessToken,
		Description: token.Description,
		IssuedAt:    timestamppb.New(token.IssuedAt),
		ExpiresAt:   timestamppb.New(token.ExpiresAt),
	}
}

func convertUserWebhook(webhook *model.UserWebhook) *modelpb.UserWebhook {
	return &modelpb.UserWebhook{
		Name:        webhook.Name,
		Url:         webhook.URL,
		DisplayName: webhook.DisplayName,
		CreateTime:  timestamppb.New(webhook.CreateTime),
		UpdateTime:  timestamppb.New(webhook.UpdateTime),
	}
}

func convertUserSettingGeneralSetting(setting *model.GeneralUserSetting) *modelpb.UserSetting_GeneralSetting {
	return &modelpb.UserSetting_GeneralSetting{
		Locale:         setting.Locale,
		Appearance:     setting.Appearance,
		MemoVisibility: setting.MemoVisibility,
		Theme:          setting.Theme,
	}
}

func convertUserSettingToProto(setting *model.UserSetting) *modelpb.UserSetting {
	info := &modelpb.UserSetting{Name: setting.Key}
	if setting.Value.GeneralUserSetting != nil {
		info.Value = &modelpb.UserSetting_GeneralSetting_{
			GeneralSetting: convertUserSettingGeneralSetting(setting.Value.GeneralUserSetting),
		}
	} else if setting.Value.SessionsUserSetting != nil {
		info.Value = &modelpb.UserSetting_SessionsSetting_{
			SessionsSetting: &modelpb.UserSetting_SessionsSetting{
				Sessions: lo.Map(setting.Value.SessionsUserSetting.Sessions, func(session *model.UserSession, _ int) *modelpb.UserSession {
					return convertUserSessionToProto(session)
				}),
			},
		}
	} else if setting.Value.AccessTokensUserSetting != nil {
		info.Value = &modelpb.UserSetting_AccessTokensSetting_{
			AccessTokensSetting: &modelpb.UserSetting_AccessTokensSetting{
				AccessTokens: lo.Map(setting.Value.AccessTokensUserSetting.AccessTokens, func(token *model.UserAccessToken, _ int) *modelpb.UserAccessToken {
					return convertUserAccessToken(token)
				}),
			},
		}
	} else if setting.Value.WebhooksUserSetting != nil {
		info.Value = &modelpb.UserSetting_WebhooksSetting_{
			WebhooksSetting: &modelpb.UserSetting_WebhooksSetting{
				Webhooks: lo.Map(setting.Value.WebhooksUserSetting.Webhooks, func(webhook *model.UserWebhook, _ int) *modelpb.UserWebhook {
					return convertUserWebhook(webhook)
				}),
			},
		}
	}
	return info
}
