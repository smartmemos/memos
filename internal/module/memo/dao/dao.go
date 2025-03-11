package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/memo"
)

type Dao struct {
}

func New(i do.Injector) (memo.DAO, error) {
	return &Dao{}, nil
}
