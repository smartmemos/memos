//go:generate mockgen -source $GOFILE -destination ./dao_mock.go -package $GOPACKAGE
package workspace

import (
	"context"

	"github.com/smartmemos/memos/internal/module/workspace/model"
)

type DAO interface {
	CreateProfile(ctx context.Context, m *model.Profile) error
	UpdateProfiles(ctx context.Context, filter *model.FindProfileFilter, update map[string]any) (int64, error)
	UpdateProfile(ctx context.Context, m *model.Profile, update map[string]any) error
	CountProfiles(ctx context.Context, filter *model.FindProfileFilter) (int64, error)
	FindProfiles(ctx context.Context, filter *model.FindProfileFilter) ([]*model.Profile, error)
	FindProfileByID(ctx context.Context, id int64) (*model.Profile, error)
	FindProfile(ctx context.Context, filter *model.FindProfileFilter) (*model.Profile, error)
}
