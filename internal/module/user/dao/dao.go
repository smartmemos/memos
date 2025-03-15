package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/user"
)

type Dao struct {
}

func New(i do.Injector) (user.DAO, error) {
	return &Dao{}, nil
}
