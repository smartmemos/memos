package dao

import (
	"github.com/samber/do/v2"

	"github.com/smartmemos/memos/internal/module/system"
)

type Dao struct {
}

func New(i do.Injector) (system.DAO, error) {
	return &Dao{}, nil
}
