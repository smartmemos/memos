package service

import (
	"context"

	usermd "github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/module/workspace/model"
)

func (s *Service) GetProfile(ctx context.Context, _ *model.GetProfileRequest) (profile *model.Profile, err error) {
	profile = &model.Profile{}
	user, err := s.userDao.FindUser(ctx, &usermd.FindUserFilter{
		Role: usermd.RoleHost,
	})
	if err != nil {
		return
	}
	profile.Owner = user.Username
	return
}
