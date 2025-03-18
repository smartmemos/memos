package service

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	memomodel "github.com/smartmemos/memos/internal/module/memo/model"
	"github.com/smartmemos/memos/internal/module/user/model"
	wsmodel "github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
)

func (s *Service) ListAllUserStats(ctx context.Context, req *model.ListAllUserStatsRequest) (stats *model.Stats, err error) {
	setting, err := s.wsDao.FindSetting(ctx, &wsmodel.FindSettingFilter{Name: string(wsmodel.SettingKeyMemoRelated)})
	if err != nil {
		return
	}
	var value wsmodel.MemoRelatedSetting
	if err = json.Unmarshal(setting.Value.RawMessage, &value); err != nil {
		return nil, errors.Wrap(err, "failed to get workspace memo related setting")
	}

	memoFilter := &memomodel.FindMemoFilter{
		ExcludeComments: true,
		ExcludeContent:  true,
		Status:          "normal",
	}
	userID, err := grpc_util.GetUserID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %v", err)
	}
	if userID == 0 {
		memoFilter.VisibilityList = []memomodel.Visibility{memomodel.Public}
	} else {
		if memoFilter.CreatorID == 0 {
			// internalFilter := fmt.Sprintf(`creator_id == %d || visibility in ["PUBLIC", "Protected"]`, currentUser.ID)
			// if memoFind.Filter != nil {
			// 	filter := fmt.Sprintf("(%s) && (%s)", *memoFind.Filter, internalFilter)
			// 	memoFind.Filter = &filter
			// } else {
			// 	memoFind.Filter = &internalFilter
			// }
		} else if memoFilter.CreatorID != userID {
			memoFilter.VisibilityList = []memomodel.Visibility{memomodel.Public, memomodel.Protected}
		}
	}

	memos, err := s.memoDao.FindMemos(ctx, memoFilter)
	if err != nil {
		return
	}
	for _, memo := range memos {
		logrus.Info(memo)
	}
	return
}
