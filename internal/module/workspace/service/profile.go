package service

import (
	"context"

	usermd "github.com/smartmemos/memos/internal/module/user/model"
	"github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) GetProfile(ctx context.Context, _ *model.GetProfileRequest) (profile *model.Profile, err error) {
	profile = &model.Profile{
		Mode:        "prod",
		Version:     "1.0",
		InstanceUrl: "http://localhost:8888",
	}
	user, err := s.userDao.FindUser(ctx, &usermd.FindUserFilter{
		Role: db.Eq(usermd.RoleHost),
	})
	if err != nil {
		return
	}
	profile.Owner = user.Username
	return
}
