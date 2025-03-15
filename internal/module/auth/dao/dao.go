package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/auth"
)

type Dao struct {
}

func New(i do.Injector) (auth.DAO, error) {
	return &Dao{}, nil
}
