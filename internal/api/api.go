package api

import (
	"github.com/samber/do/v2"

	v1 "github.com/smartmemos/memos/internal/api/v1"
)

func Init(i do.Injector) {
	do.Provide(i, v1.NewAuthService)
	do.Provide(i, v1.NewMemoService)
	do.Provide(i, v1.NewUserService)
	do.Provide(i, v1.NewWorkspaceService)
}
