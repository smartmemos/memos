package api

import (
	"github.com/samber/do/v2"

	v1 "github.com/smartmemos/memos/internal/api/v1"
	v2 "github.com/smartmemos/memos/internal/api/v2"
)

func Init(i do.Injector) {
	// api jv1
	do.Provide(i, v1.NewAuthService)
	do.Provide(i, v1.NewMemoService)
	do.Provide(i, v1.NewUserService)
	do.Provide(i, v1.NewWorkspaceService)

	// api v2
	do.Provide(i, v2.NewAuthService)

}
