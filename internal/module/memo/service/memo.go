package service

import (
	"context"
	"slices"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"

	"github.com/smartmemos/memos/internal/module/memo/model"
	usermd "github.com/smartmemos/memos/internal/module/user/model"
	wsmd "github.com/smartmemos/memos/internal/module/workspace/model"
	"github.com/smartmemos/memos/internal/pkg/grpc_util"
)

func (s *Service) CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (memo *model.Memo, err error) {
	setting, err := s.wsDao.FindMemoRelatedSetting(ctx)
	if err != nil {
		return
	}
	if setting.DisallowPublicVisibility && req.Visibility == model.Public {
		err = errors.New("disable public memos system setting is enabled")
		return
	}
	if contentLimit := int(setting.ContentLengthLimit); len(req.Content) > contentLimit {
		err = errors.Errorf("content too long (max %d characters)", contentLimit)
		return
	}
	_, err = s.getMemoPayload(req.Content)
	if err != nil {
		return
	}
	memo = &model.Memo{
		Content:      req.Content,
		Visibility:   req.Visibility,
		ParentID:     0,
		RelationType: "",
		Status:       "normal",
		Pinned:       true,
	}
	if err = s.dao.CreateMemo(ctx, memo); err != nil {
		return
	}

	// if len(req.Resources) > 0 {
	// 	_, err := s.SetMemoResources(ctx, &v1pb.SetMemoResourcesRequest{
	// 		Name:      fmt.Sprintf("%s%s", MemoNamePrefix, memo.UID),
	// 		Resources: req.Memo.Resources,
	// 	})
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "failed to set memo resources")
	// 	}
	// }
	// if len(req.Relations) > 0 {
	// 	_, err := s.SetMemoRelations(ctx, &v1pb.SetMemoRelationsRequest{
	// 		Name:      fmt.Sprintf("%s%s", MemoNamePrefix, memo.UID),
	// 		Relations: request.Memo.Relations,
	// 	})
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "failed to set memo relations")
	// 	}
	// }

	return
}

func (s *Service) getMemoPayload(content string) (payload *model.MemoPayload, err error) {
	nodes, err := parser.Parse(tokenizer.Tokenize(content))
	if err != nil {
		err = errors.Wrap(err, "failed to parse content")
		return
	}

	payload = &model.MemoPayload{}
	tags := []string{}
	property := &model.MemoPayloadProperty{}
	TraverseASTNodes(nodes, func(node ast.Node) {
		switch n := node.(type) {
		case *ast.Tag:
			tag := n.Content
			if !slices.Contains(tags, tag) {
				tags = append(tags, tag)
			}
		case *ast.Link, *ast.AutoLink:
			property.HasLink = true
		case *ast.TaskListItem:
			property.HasTaskList = true
			if !n.Complete {
				property.HasIncompleteTasks = true
			}
		case *ast.Code, *ast.CodeBlock:
			property.HasCode = true
		case *ast.EmbeddedContent:
			// TODO: validate references.
			property.References = append(property.References, n.ResourceName)
		}
	})
	payload.Tags = tags
	payload.Property = property
	return
}

func TraverseASTNodes(nodes []ast.Node, fn func(ast.Node)) {
	for _, node := range nodes {
		fn(node)
		switch n := node.(type) {
		case *ast.Paragraph:
			TraverseASTNodes(n.Children, fn)
		case *ast.Heading:
			TraverseASTNodes(n.Children, fn)
		case *ast.Blockquote:
			TraverseASTNodes(n.Children, fn)
		case *ast.List:
			TraverseASTNodes(n.Children, fn)
		case *ast.OrderedListItem:
			TraverseASTNodes(n.Children, fn)
		case *ast.UnorderedListItem:
			TraverseASTNodes(n.Children, fn)
		case *ast.TaskListItem:
			TraverseASTNodes(n.Children, fn)
		case *ast.Bold:
			TraverseASTNodes(n.Children, fn)
		}
	}
}

func (s *Service) UpdateMemo(ctx context.Context, req *model.UpdateMemoRequest) (memoInfo *model.MemoInfo, err error) {
	userId, err := grpc_util.GetUserID(ctx)
	if err != nil {
		return
	}
	user, err := s.userDao.FindUserByID(ctx, userId)
	if err != nil {
		return
	}
	memo, err := s.dao.FindMemoByID(ctx, req.ID)
	if err != nil {
		return
	}
	// Only the creator or admin can update the memo.
	if memo.CreatorID != user.ID && !s.isSuperUser(user) {
		err = errors.New("permission denied")
		return
	}

	update := make(map[string]any)
	if lo.Contains(req.Paths, "content") {
		var setting *wsmd.MemoRelatedSetting
		setting, err = s.wsDao.FindMemoRelatedSetting(ctx)
		if err != nil {
			return
		}
		if contentLimit := int(setting.ContentLengthLimit); len(req.Content) > contentLimit {
			err = errors.Errorf("content too long (max %d characters)", contentLimit)
			return
		}
		_, err = s.getMemoPayload(req.Content)
		if err != nil {
			return
		}
	}

	if lo.Contains(req.Paths, "visibility") {
	}
	if lo.Contains(req.Paths, "pinned") {
		update["pinned"] = req.Pinned
	}
	if lo.Contains(req.Paths, "state") {
	}
	if lo.Contains(req.Paths, "create_time") {
	}
	if lo.Contains(req.Paths, "update_time") {
	}
	if lo.Contains(req.Paths, "display_time") {
	}
	if lo.Contains(req.Paths, "location") {
	}
	if lo.Contains(req.Paths, "resources") {
	}
	if lo.Contains(req.Paths, "relations") {
	}

	err = s.dao.UpdateMemo(ctx, memo, update)
	if err != nil {
		return
	}
	memoInfo = &model.MemoInfo{
		Memo: memo,
	}
	return
}

func (s *Service) DeleteMemo(ctx context.Context, id int64) (err error) {
	userId, err := grpc_util.GetUserID(ctx)
	if err != nil {
		return
	}
	user, err := s.userDao.FindUserByID(ctx, userId)
	if err != nil {
		return
	}
	memo, err := s.dao.FindMemoByID(ctx, id)
	if err != nil {
		return
	}
	// Only the creator or admin can update the memo.
	if memo.CreatorID != userId && !s.isSuperUser(user) {
		err = errors.New("permission denied")
		return
	}
	_, err = s.dao.DeleteMemos(ctx, &model.FindMemoFilter{ID: id})
	if err != nil {
		return
	}
	return
}

func (s *Service) isSuperUser(user *usermd.User) bool {
	return user.Role == usermd.RoleAdmin || user.Role == usermd.RoleHost
}

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (list []*model.MemoInfo, err error) {
	memos, err := s.dao.FindMemos(ctx, &model.FindMemoFilter{})
	if err != nil {
		return
	}
	var pids []int64
	for _, memo := range memos {
		if memo.ParentID > 0 {
			pids = append(pids, memo.ParentID)
		}
	}
	var memosMap map[int64]*model.Memo
	if len(pids) > 0 {
		var parentList []*model.Memo
		parentList, err = s.dao.FindMemos(ctx, &model.FindMemoFilter{ParentIDs: pids})
		if err != nil {
			return
		}
		memosMap = lo.Associate(parentList, func(item *model.Memo) (int64, *model.Memo) { return item.ID, item })
	}
	for _, memo := range memos {
		if v, ok := memosMap[memo.ID]; ok {
			logrus.Info(v)
		}
		item := &model.MemoInfo{
			Memo: memo,
		}
		list = append(list, item)
	}
	return
}
