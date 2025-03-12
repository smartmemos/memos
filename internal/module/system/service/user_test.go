package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/samber/do/v2"
	"github.com/smartmemos/memos/internal/module/system"
	"github.com/smartmemos/memos/internal/module/system/model"
)

func TestService_CreateUser(t *testing.T) {
	ctx := context.TODO()

	ctl := gomock.NewController(t)
	daoMock := system.NewMockDAO(ctl)

	daoMock.
		EXPECT().
		CreateUser(ctx, gomock.Any()).
		Return(nil)

	do.Override(nil, func(i do.Injector) (system.DAO, error) {
		return daoMock, nil
	})

	s := do.MustInvoke[*Service](nil)
	user, err := s.CreateUser(ctx, &model.CreateUserRequest{
		Username: "User-123456",
		Password: "123456",
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Logf("user: %v", user)
}
