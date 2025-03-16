package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/user"
	"github.com/smartmemos/memos/internal/module/workspace"
	"github.com/smartmemos/memos/internal/module/workspace/dao"
)

type Service struct {
	dao     workspace.DAO
	userDao user.DAO
}

func New(i do.Injector) (workspace.Service, error) {
	return &Service{
		dao:     do.MustInvoke[workspace.DAO](i),
		userDao: do.MustInvoke[user.DAO](i),
	}, nil
}

func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
