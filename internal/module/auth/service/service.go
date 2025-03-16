package service

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/auth"
	"github.com/smartmemos/memos/internal/module/auth/dao"
	"github.com/smartmemos/memos/internal/module/user"
)

type Service struct {
	dao     auth.DAO
	userDao user.DAO
}

func New(i do.Injector) (auth.Service, error) {
	return &Service{
		dao:     do.MustInvoke[auth.DAO](i),
		userDao: do.MustInvoke[user.DAO](i),
	}, nil
}

func Init(i do.Injector) {
	do.Provide(i, dao.New)
	do.Provide(i, New)
}
