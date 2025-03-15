package module

import (
	"github.com/samber/do/v2"

	authimpl "github.com/smartmemos/memos/internal/module/auth/service"
	memoimpl "github.com/smartmemos/memos/internal/module/memo/service"
	userimpl "github.com/smartmemos/memos/internal/module/user/service"
	workspaceimpl "github.com/smartmemos/memos/internal/module/workspace/service"
)

func Init(i do.Injector) {
	authimpl.Init(i)
	memoimpl.Init(i)
	userimpl.Init(i)
	workspaceimpl.Init(i)
}
