package api

import (
	"github.com/samber/do/v2"

	apiv2 "github.com/smartmemos/memos/internal/api/v2"
)

func Init(i do.Injector) {
	// api v2
	do.Provide(i, apiv2.NewAuthService)
	do.Provide(i, apiv2.NewUserService)
	do.Provide(i, apiv2.NewInboxService)
	do.Provide(i, apiv2.NewMemoService)
	do.Provide(i, apiv2.NewWorkspaceService)
	do.Provide(i, apiv2.NewMarkdownService)
	do.Provide(i, apiv2.NewWebhookService)
}
