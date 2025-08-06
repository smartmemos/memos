package service

import (
	"context"
	"slices"

	"github.com/lithammer/shortuuid/v4"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"

	"github.com/smartmemos/memos/internal/memos/model"
	"github.com/smartmemos/memos/internal/pkg/db"
)

func (s *Service) CreateMemo(ctx context.Context, req *model.CreateMemoRequest) (memo *model.Memo, err error) {
	memo = &model.Memo{
		UID:        shortuuid.New(),
		CreatorID:  req.UserID,
		Content:    req.Content,
		Visibility: req.Visibility,
	}
	memoRelatedSetting, err := s.GetMemoRelatedSetting(ctx)
	if err != nil {
		err = errors.New("failed to get workspace memo related setting")
		return
	}
	if memoRelatedSetting.DisallowPublicVisibility && memo.Visibility == model.Public {
		err = errors.New("disable public memos system setting is enabled")
		return
	}
	if contentLimit := memoRelatedSetting.ContentLengthLimit; len(memo.Content) > contentLimit {
		err = errors.Errorf("content too long (max %d characters)", contentLimit)
		return
	}
	if err = s.RebuildMemoPayload(memo); err != nil {
		err = errors.Errorf("failed to rebuild memo payload: %v", err)
		return
	}
	if req.Location != nil {
		memo.Payload.Location = req.Location
	}
	err = s.dao.CreateMemo(ctx, memo)
	if err != nil {
		err = errors.Wrap(err, "failed to create memo")
		return
	}

	// if len(req.Attachments) > 0 {
	// 	_, err := s.SetMemoAttachments(ctx, &v1pb.SetMemoAttachmentsRequest{
	// 		Name:        fmt.Sprintf("%s%s", MemoNamePrefix, memo.UID),
	// 		Attachments: request.Memo.Attachments,
	// 	})
	// 	if err != nil {
	// 		return nil, errors.Wrap(err, "failed to set memo attachments")
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

func (s *Service) RebuildMemoPayload(memo *model.Memo) error {
	nodes, err := parser.Parse(tokenizer.Tokenize(memo.Content))
	if err != nil {
		return errors.Wrap(err, "failed to parse content")
	}

	if memo.Payload == nil {
		memo.Payload = &model.MemoPayload{}
	}
	tags := []string{}
	property := &model.MemoPayloadProperty{}
	s.TraverseASTNodes(nodes, func(node ast.Node) {
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
		case *ast.CodeBlock:
			property.HasCode = true
		case *ast.EmbeddedContent:
			// TODO: validate references.
			property.References = append(property.References, n.ResourceName)
		}
	})
	memo.Payload.Tags = tags
	memo.Payload.Property = property
	return nil
}

func (s *Service) TraverseASTNodes(nodes []ast.Node, fn func(ast.Node)) {
	for _, node := range nodes {
		fn(node)
		switch n := node.(type) {
		case *ast.Paragraph:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.Heading:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.Blockquote:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.List:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.OrderedListItem:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.UnorderedListItem:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.TaskListItem:
			s.TraverseASTNodes(n.Children, fn)
		case *ast.Bold:
			s.TraverseASTNodes(n.Children, fn)
		}
	}
}

func (s *Service) ListMemos(ctx context.Context, req *model.ListMemosRequest) (total int64, list []*model.Memo, err error) {
	filter := &model.FindMemoFilter{}

	total, err = s.dao.CountMemos(ctx, filter)
	if err != nil {
		return
	}

	memos, err := s.dao.FindMemos(ctx, filter)
	if err != nil {
		return
	}
	pids := lo.FlatMap(memos, func(item *model.Memo, _ int) []int64 {
		return lo.If(item.ParentID > 0, []int64{item.ParentID}).Else(nil)
	})
	var memosMap map[int64]*model.Memo
	if len(pids) > 0 {
		var parentList []*model.Memo
		parentList, err = s.dao.FindMemos(ctx, &model.FindMemoFilter{ParentIDs: db.In(pids)})
		if err != nil {
			return
		}
		memosMap = lo.Associate(parentList, func(item *model.Memo) (int64, *model.Memo) {
			return item.ID, item
		})
	}

	for _, memo := range memos {
		if v, ok := memosMap[memo.ID]; ok {
			memo.ParentID = v.ID
		}
		list = append(list, memo)
	}
	return
}

func (s *Service) GetMemo(ctx context.Context, req *model.GetMemoRequest) (memo *model.Memo, err error) {

	return
}
