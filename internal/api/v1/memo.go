package v1

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/samber/do/v2"
	"github.com/samber/lo"
	"github.com/usememos/gomark/ast"
	"github.com/usememos/gomark/parser"
	"github.com/usememos/gomark/parser/tokenizer"
	"github.com/usememos/gomark/renderer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartmemos/memos/internal/module/memo"
	"github.com/smartmemos/memos/internal/module/memo/model"
	"github.com/smartmemos/memos/internal/module/workspace"
	v1pb "github.com/smartmemos/memos/internal/proto/api/v1"
	"github.com/smartmemos/memos/internal/proto/model/common"
	mdpb "github.com/smartmemos/memos/internal/proto/model/markdown"
	memopb "github.com/smartmemos/memos/internal/proto/model/memo"
)

type MemoService struct {
	v1pb.UnimplementedMemoServiceServer
	memoService      memo.Service
	workspaceService workspace.Service
}

func NewMemoService(i do.Injector) (*MemoService, error) {
	return &MemoService{
		memoService: do.MustInvoke[memo.Service](i),
	}, nil
}

func (s *MemoService) CreateMemo(ctx context.Context, req *v1pb.CreateMemoRequest) (resp *memopb.Memo, err error) {
	var relations []model.MemoRelation
	for _, item := range req.Memo.Relations {
		relations = append(relations, model.MemoRelation{
			Type: model.RelationType(memopb.MemoRelation_Type_name[int32(item.Type)]),
			Memo: model.RelationMemo{
				Uid:     item.Memo.Uid,
				Name:    item.Memo.Name,
				Snippet: item.Memo.Snippet,
			},
			RelatedMemo: model.RelationMemo{
				Uid:     item.RelatedMemo.Uid,
				Name:    item.RelatedMemo.Name,
				Snippet: item.RelatedMemo.Snippet,
			},
		})
	}
	memo, err := s.memoService.CreateMemo(ctx, &model.CreateMemoRequest{
		Content:    req.Memo.Content,
		Visibility: convertFromProtoVisibility(req.Memo.Visibility),
		Relations:  relations,
	})
	if err != nil {
		return
	}
	resp = convertMemoToProto(memo)
	return
}

func convertMemoToProto(memo *model.Memo) *memopb.Memo {
	resp := &memopb.Memo{
		Name: fmt.Sprintf("%d", memo.ID),
		// State:       convertStateFromStore(memo.RowStatus),
		// Creator:     fmt.Sprintf("%s%d", UserNamePrefix, memo.CreatorID),
		CreateTime: timestamppb.New(memo.CreatedAt),
		UpdateTime: timestamppb.New(memo.UpdatedAt),
		// DisplayTime: timestamppb.New(time.Unix(displayTs, 0)),
		Content:    memo.Content,
		Visibility: convertVisibilityToProto(memo.Visibility),
		// Pinned:     memo.Pinned,
	}
	return resp
}

func convertFromProtoVisibility(v memopb.Visibility) model.Visibility {
	return model.Visibility(memopb.Visibility_name[int32(v)])
}

func convertVisibilityToProto(v model.Visibility) memopb.Visibility {
	return memopb.Visibility(memopb.Visibility_value[string(v)])
}

func (s *MemoService) ListMemos(ctx context.Context, req *v1pb.ListMemosRequest) (resp *v1pb.ListMemosResponse, err error) {
	var pageSize, page int
	if req.PageToken != "" {
		page, pageSize, err = parsePageToken(req.PageToken)
		if err != nil {
			return
		}
	} else {
		pageSize = int(req.PageSize)
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	memos, err := s.memoService.ListMemos(ctx, &model.ListMemosRequest{
		PageSize: pageSize + 1,
		Page:     page,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "failed to list memos: %v", err)
		return
	}
	nextPageToken := ""
	if len(memos) > pageSize {
		memos = memos[:pageSize]
		nextPageToken, err = getPageToken(pageSize, page+1)
		if err != nil {
			err = status.Errorf(codes.Internal, "failed to get next page token, error: %v", err)
			return
		}
	}
	var list []*memopb.Memo
	for _, memo := range memos {
		item, err := s.convertMemoToProto(ctx, memo)
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert memo")
		}
		list = append(list, item)
	}

	resp = &v1pb.ListMemosResponse{
		Memos:         list,
		NextPageToken: nextPageToken,
	}
	return
}

func (s *MemoService) convertMemoToProto(ctx context.Context, memo *model.MemoInfo) (item *memopb.Memo, err error) {
	name := fmt.Sprintf("%s%d", MemoNamePrefix, memo.ID)
	snippet, _ := getMemoContentSnippet(memo.Content)

	nodes, err := parser.Parse(tokenizer.Tokenize(memo.Content))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse content")
	}

	item = &memopb.Memo{
		Name:        name,
		State:       common.State(common.State_value[memo.Status]),
		Creator:     fmt.Sprintf("%s%d", UserNamePrefix, memo.CreatorID),
		CreateTime:  timestamppb.New(memo.CreatedAt),
		UpdateTime:  timestamppb.New(memo.UpdatedAt),
		DisplayTime: timestamppb.New(memo.UpdatedAt),
		Content:     memo.Content,
		Visibility:  convertVisibilityToProto(memo.Visibility),
		Pinned:      memo.Pinned,
		Snippet:     snippet,
		Nodes:       convertFromASTNodes(nodes),
	}
	return
}

func convertFromASTNode(rawNode ast.Node) *mdpb.Node {
	node := &mdpb.Node{
		Type: mdpb.NodeType(mdpb.NodeType_value[string(rawNode.Type())]),
	}

	switch n := rawNode.(type) {
	case *ast.LineBreak:
		node.Node = &mdpb.Node_LineBreakNode{}
	case *ast.Paragraph:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_ParagraphNode{ParagraphNode: &mdpb.ParagraphNode{Children: children}}
	case *ast.CodeBlock:
		node.Node = &mdpb.Node_CodeBlockNode{CodeBlockNode: &mdpb.CodeBlockNode{Language: n.Language, Content: n.Content}}
	case *ast.Heading:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_HeadingNode{HeadingNode: &mdpb.HeadingNode{Level: int32(n.Level), Children: children}}
	case *ast.HorizontalRule:
		node.Node = &mdpb.Node_HorizontalRuleNode{HorizontalRuleNode: &mdpb.HorizontalRuleNode{Symbol: n.Symbol}}
	case *ast.Blockquote:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_BlockquoteNode{BlockquoteNode: &mdpb.BlockquoteNode{Children: children}}
	case *ast.List:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_ListNode{ListNode: &mdpb.ListNode{Kind: convertListKindFromASTNode(n.Kind), Indent: int32(n.Indent), Children: children}}
	case *ast.OrderedListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_OrderedListItemNode{OrderedListItemNode: &mdpb.OrderedListItemNode{Number: n.Number, Indent: int32(n.Indent), Children: children}}
	case *ast.UnorderedListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_UnorderedListItemNode{UnorderedListItemNode: &mdpb.UnorderedListItemNode{Symbol: n.Symbol, Indent: int32(n.Indent), Children: children}}
	case *ast.TaskListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &mdpb.Node_TaskListItemNode{TaskListItemNode: &mdpb.TaskListItemNode{Symbol: n.Symbol, Indent: int32(n.Indent), Complete: n.Complete, Children: children}}
	case *ast.MathBlock:
		node.Node = &mdpb.Node_MathBlockNode{MathBlockNode: &mdpb.MathBlockNode{Content: n.Content}}
	case *ast.Table:
		node.Node = &mdpb.Node_TableNode{TableNode: convertTableFromASTNode(n)}
	case *ast.EmbeddedContent:
		node.Node = &mdpb.Node_EmbeddedContentNode{EmbeddedContentNode: &mdpb.EmbeddedContentNode{ResourceName: n.ResourceName, Params: n.Params}}
	case *ast.Text:
		node.Node = &mdpb.Node_TextNode{TextNode: &mdpb.TextNode{Content: n.Content}}
	case *ast.Bold:
		node.Node = &mdpb.Node_BoldNode{BoldNode: &mdpb.BoldNode{Symbol: n.Symbol, Children: convertFromASTNodes(n.Children)}}
	case *ast.Italic:
		node.Node = &mdpb.Node_ItalicNode{ItalicNode: &mdpb.ItalicNode{Symbol: n.Symbol, Content: n.Content}}
	case *ast.BoldItalic:
		node.Node = &mdpb.Node_BoldItalicNode{BoldItalicNode: &mdpb.BoldItalicNode{Symbol: n.Symbol, Content: n.Content}}
	case *ast.Code:
		node.Node = &mdpb.Node_CodeNode{CodeNode: &mdpb.CodeNode{Content: n.Content}}
	case *ast.Image:
		node.Node = &mdpb.Node_ImageNode{ImageNode: &mdpb.ImageNode{AltText: n.AltText, Url: n.URL}}
	case *ast.Link:
		node.Node = &mdpb.Node_LinkNode{LinkNode: &mdpb.LinkNode{Content: convertFromASTNodes(n.Content), Url: n.URL}}
	case *ast.AutoLink:
		node.Node = &mdpb.Node_AutoLinkNode{AutoLinkNode: &mdpb.AutoLinkNode{Url: n.URL, IsRawText: n.IsRawText}}
	case *ast.Tag:
		node.Node = &mdpb.Node_TagNode{TagNode: &mdpb.TagNode{Content: n.Content}}
	case *ast.Strikethrough:
		node.Node = &mdpb.Node_StrikethroughNode{StrikethroughNode: &mdpb.StrikethroughNode{Content: n.Content}}
	case *ast.EscapingCharacter:
		node.Node = &mdpb.Node_EscapingCharacterNode{EscapingCharacterNode: &mdpb.EscapingCharacterNode{Symbol: n.Symbol}}
	case *ast.Math:
		node.Node = &mdpb.Node_MathNode{MathNode: &mdpb.MathNode{Content: n.Content}}
	case *ast.Highlight:
		node.Node = &mdpb.Node_HighlightNode{HighlightNode: &mdpb.HighlightNode{Content: n.Content}}
	case *ast.Subscript:
		node.Node = &mdpb.Node_SubscriptNode{SubscriptNode: &mdpb.SubscriptNode{Content: n.Content}}
	case *ast.Superscript:
		node.Node = &mdpb.Node_SuperscriptNode{SuperscriptNode: &mdpb.SuperscriptNode{Content: n.Content}}
	case *ast.ReferencedContent:
		node.Node = &mdpb.Node_ReferencedContentNode{ReferencedContentNode: &mdpb.ReferencedContentNode{ResourceName: n.ResourceName, Params: n.Params}}
	case *ast.Spoiler:
		node.Node = &mdpb.Node_SpoilerNode{SpoilerNode: &mdpb.SpoilerNode{Content: n.Content}}
	case *ast.HTMLElement:
		node.Node = &mdpb.Node_HtmlElementNode{HtmlElementNode: &mdpb.HTMLElementNode{TagName: n.TagName, Attributes: n.Attributes}}
	default:
		node.Node = &mdpb.Node_TextNode{TextNode: &mdpb.TextNode{}}
	}
	return node
}

func convertTableFromASTNode(node *ast.Table) *mdpb.TableNode {
	table := &mdpb.TableNode{
		Header:    convertFromASTNodes(node.Header),
		Delimiter: node.Delimiter,
	}
	for _, row := range node.Rows {
		table.Rows = append(table.Rows, &mdpb.TableNode_Row{Cells: convertFromASTNodes(row)})
	}
	return table
}

func convertListKindFromASTNode(node ast.ListKind) mdpb.ListNode_Kind {
	switch node {
	case ast.OrderedList:
		return mdpb.ListNode_ORDERED
	case ast.UnorderedList:
		return mdpb.ListNode_UNORDERED
	case ast.DescrpitionList:
		return mdpb.ListNode_DESCRIPTION
	default:
		return mdpb.ListNode_KIND_UNSPECIFIED
	}
}

func convertFromASTNodes(rawNodes []ast.Node) []*mdpb.Node {
	nodes := []*mdpb.Node{}
	for _, rawNode := range rawNodes {
		node := convertFromASTNode(rawNode)
		nodes = append(nodes, node)
	}
	return nodes
}

func getMemoContentSnippet(content string) (string, error) {
	nodes, err := parser.Parse(tokenizer.Tokenize(content))
	if err != nil {
		return "", errors.Wrap(err, "failed to parse content")
	}

	plainText := renderer.NewStringRenderer().Render(nodes)
	if len(plainText) > 64 {
		return lo.Substring(plainText, 0, 64) + "...", nil
	}
	return plainText, nil
}
