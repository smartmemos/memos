package v2

import (
	"github.com/samber/do/v2"
	"github.com/usememos/gomark/ast"

	"github.com/smartmemos/memos/internal/memos"
	v2pb "github.com/smartmemos/memos/internal/proto/api/v2"
	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

type MarkdownService struct {
	v2pb.UnimplementedMarkdownServiceHandler
	memosService memos.Service
}

func NewMarkdownService(i do.Injector) (*MarkdownService, error) {
	return &MarkdownService{
		memosService: do.MustInvoke[memos.Service](i),
	}, nil
}

func convertFromASTNode(rawNode ast.Node) *modelpb.Node {
	node := &modelpb.Node{
		Type: modelpb.NodeType(modelpb.NodeType_value[string(rawNode.Type())]),
	}

	switch n := rawNode.(type) {
	case *ast.LineBreak:
		node.Node = &modelpb.Node_LineBreakNode{}
	case *ast.Paragraph:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_ParagraphNode{ParagraphNode: &modelpb.ParagraphNode{Children: children}}
	case *ast.CodeBlock:
		node.Node = &modelpb.Node_CodeBlockNode{CodeBlockNode: &modelpb.CodeBlockNode{Language: n.Language, Content: n.Content}}
	case *ast.Heading:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_HeadingNode{HeadingNode: &modelpb.HeadingNode{Level: int32(n.Level), Children: children}}
	case *ast.HorizontalRule:
		node.Node = &modelpb.Node_HorizontalRuleNode{HorizontalRuleNode: &modelpb.HorizontalRuleNode{Symbol: n.Symbol}}
	case *ast.Blockquote:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_BlockquoteNode{BlockquoteNode: &modelpb.BlockquoteNode{Children: children}}
	case *ast.List:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_ListNode{ListNode: &modelpb.ListNode{Kind: convertListKindFromASTNode(n.Kind), Indent: int32(n.Indent), Children: children}}
	case *ast.OrderedListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_OrderedListItemNode{OrderedListItemNode: &modelpb.OrderedListItemNode{Number: n.Number, Indent: int32(n.Indent), Children: children}}
	case *ast.UnorderedListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_UnorderedListItemNode{UnorderedListItemNode: &modelpb.UnorderedListItemNode{Symbol: n.Symbol, Indent: int32(n.Indent), Children: children}}
	case *ast.TaskListItem:
		children := convertFromASTNodes(n.Children)
		node.Node = &modelpb.Node_TaskListItemNode{TaskListItemNode: &modelpb.TaskListItemNode{Symbol: n.Symbol, Indent: int32(n.Indent), Complete: n.Complete, Children: children}}
	case *ast.MathBlock:
		node.Node = &modelpb.Node_MathBlockNode{MathBlockNode: &modelpb.MathBlockNode{Content: n.Content}}
	case *ast.Table:
		node.Node = &modelpb.Node_TableNode{TableNode: convertTableFromASTNode(n)}
	case *ast.EmbeddedContent:
		node.Node = &modelpb.Node_EmbeddedContentNode{EmbeddedContentNode: &modelpb.EmbeddedContentNode{ResourceName: n.ResourceName, Params: n.Params}}
	case *ast.Text:
		node.Node = &modelpb.Node_TextNode{TextNode: &modelpb.TextNode{Content: n.Content}}
	case *ast.Bold:
		node.Node = &modelpb.Node_BoldNode{BoldNode: &modelpb.BoldNode{Symbol: n.Symbol, Children: convertFromASTNodes(n.Children)}}
	case *ast.Italic:
		node.Node = &modelpb.Node_ItalicNode{ItalicNode: &modelpb.ItalicNode{Symbol: n.Symbol, Children: convertFromASTNodes(n.Children)}}
	case *ast.BoldItalic:
		node.Node = &modelpb.Node_BoldItalicNode{BoldItalicNode: &modelpb.BoldItalicNode{Symbol: n.Symbol, Content: n.Content}}
	case *ast.Code:
		node.Node = &modelpb.Node_CodeNode{CodeNode: &modelpb.CodeNode{Content: n.Content}}
	case *ast.Image:
		node.Node = &modelpb.Node_ImageNode{ImageNode: &modelpb.ImageNode{AltText: n.AltText, Url: n.URL}}
	case *ast.Link:
		node.Node = &modelpb.Node_LinkNode{LinkNode: &modelpb.LinkNode{Content: convertFromASTNodes(n.Content), Url: n.URL}}
	case *ast.AutoLink:
		node.Node = &modelpb.Node_AutoLinkNode{AutoLinkNode: &modelpb.AutoLinkNode{Url: n.URL, IsRawText: n.IsRawText}}
	case *ast.Tag:
		node.Node = &modelpb.Node_TagNode{TagNode: &modelpb.TagNode{Content: n.Content}}
	case *ast.Strikethrough:
		node.Node = &modelpb.Node_StrikethroughNode{StrikethroughNode: &modelpb.StrikethroughNode{Content: n.Content}}
	case *ast.EscapingCharacter:
		node.Node = &modelpb.Node_EscapingCharacterNode{EscapingCharacterNode: &modelpb.EscapingCharacterNode{Symbol: n.Symbol}}
	case *ast.Math:
		node.Node = &modelpb.Node_MathNode{MathNode: &modelpb.MathNode{Content: n.Content}}
	case *ast.Highlight:
		node.Node = &modelpb.Node_HighlightNode{HighlightNode: &modelpb.HighlightNode{Content: n.Content}}
	case *ast.Subscript:
		node.Node = &modelpb.Node_SubscriptNode{SubscriptNode: &modelpb.SubscriptNode{Content: n.Content}}
	case *ast.Superscript:
		node.Node = &modelpb.Node_SuperscriptNode{SuperscriptNode: &modelpb.SuperscriptNode{Content: n.Content}}
	case *ast.ReferencedContent:
		node.Node = &modelpb.Node_ReferencedContentNode{ReferencedContentNode: &modelpb.ReferencedContentNode{ResourceName: n.ResourceName, Params: n.Params}}
	case *ast.Spoiler:
		node.Node = &modelpb.Node_SpoilerNode{SpoilerNode: &modelpb.SpoilerNode{Content: n.Content}}
	case *ast.HTMLElement:
		node.Node = &modelpb.Node_HtmlElementNode{HtmlElementNode: &modelpb.HTMLElementNode{TagName: n.TagName, Attributes: n.Attributes}}
	default:
		node.Node = &modelpb.Node_TextNode{TextNode: &modelpb.TextNode{}}
	}
	return node
}

func convertFromASTNodes(rawNodes []ast.Node) []*modelpb.Node {
	nodes := []*modelpb.Node{}
	for _, rawNode := range rawNodes {
		node := convertFromASTNode(rawNode)
		nodes = append(nodes, node)
	}
	return nodes
}

func convertTableFromASTNode(node *ast.Table) *modelpb.TableNode {
	table := &modelpb.TableNode{
		Header:    convertFromASTNodes(node.Header),
		Delimiter: node.Delimiter,
	}
	for _, row := range node.Rows {
		table.Rows = append(table.Rows, &modelpb.TableNode_Row{Cells: convertFromASTNodes(row)})
	}
	return table
}

func convertListKindFromASTNode(node ast.ListKind) modelpb.ListNode_Kind {
	switch node {
	case ast.OrderedList:
		return modelpb.ListNode_ORDERED
	case ast.UnorderedList:
		return modelpb.ListNode_UNORDERED
	case ast.DescrpitionList:
		return modelpb.ListNode_DESCRIPTION
	default:
		return modelpb.ListNode_KIND_UNSPECIFIED
	}
}

func convertToASTNode(node *modelpb.Node) ast.Node {
	switch n := node.Node.(type) {
	case *modelpb.Node_LineBreakNode:
		return &ast.LineBreak{}
	case *modelpb.Node_ParagraphNode:
		children := convertToASTNodes(n.ParagraphNode.Children)
		return &ast.Paragraph{Children: children}
	case *modelpb.Node_CodeBlockNode:
		return &ast.CodeBlock{Language: n.CodeBlockNode.Language, Content: n.CodeBlockNode.Content}
	case *modelpb.Node_HeadingNode:
		children := convertToASTNodes(n.HeadingNode.Children)
		return &ast.Heading{Level: int(n.HeadingNode.Level), Children: children}
	case *modelpb.Node_HorizontalRuleNode:
		return &ast.HorizontalRule{Symbol: n.HorizontalRuleNode.Symbol}
	case *modelpb.Node_BlockquoteNode:
		children := convertToASTNodes(n.BlockquoteNode.Children)
		return &ast.Blockquote{Children: children}
	case *modelpb.Node_ListNode:
		children := convertToASTNodes(n.ListNode.Children)
		return &ast.List{Kind: convertListKindToASTNode(n.ListNode.Kind), Indent: int(n.ListNode.Indent), Children: children}
	case *modelpb.Node_OrderedListItemNode:
		children := convertToASTNodes(n.OrderedListItemNode.Children)
		return &ast.OrderedListItem{Number: n.OrderedListItemNode.Number, Indent: int(n.OrderedListItemNode.Indent), Children: children}
	case *modelpb.Node_UnorderedListItemNode:
		children := convertToASTNodes(n.UnorderedListItemNode.Children)
		return &ast.UnorderedListItem{Symbol: n.UnorderedListItemNode.Symbol, Indent: int(n.UnorderedListItemNode.Indent), Children: children}
	case *modelpb.Node_TaskListItemNode:
		children := convertToASTNodes(n.TaskListItemNode.Children)
		return &ast.TaskListItem{Symbol: n.TaskListItemNode.Symbol, Indent: int(n.TaskListItemNode.Indent), Complete: n.TaskListItemNode.Complete, Children: children}
	case *modelpb.Node_MathBlockNode:
		return &ast.MathBlock{Content: n.MathBlockNode.Content}
	case *modelpb.Node_TableNode:
		return convertTableToASTNode(n.TableNode)
	case *modelpb.Node_EmbeddedContentNode:
		return &ast.EmbeddedContent{ResourceName: n.EmbeddedContentNode.ResourceName, Params: n.EmbeddedContentNode.Params}
	case *modelpb.Node_TextNode:
		return &ast.Text{Content: n.TextNode.Content}
	case *modelpb.Node_BoldNode:
		return &ast.Bold{Symbol: n.BoldNode.Symbol, Children: convertToASTNodes(n.BoldNode.Children)}
	case *modelpb.Node_ItalicNode:
		return &ast.Italic{Symbol: n.ItalicNode.Symbol, Children: convertToASTNodes(n.ItalicNode.Children)}
	case *modelpb.Node_BoldItalicNode:
		return &ast.BoldItalic{Symbol: n.BoldItalicNode.Symbol, Content: n.BoldItalicNode.Content}
	case *modelpb.Node_CodeNode:
		return &ast.Code{Content: n.CodeNode.Content}
	case *modelpb.Node_ImageNode:
		return &ast.Image{AltText: n.ImageNode.AltText, URL: n.ImageNode.Url}
	case *modelpb.Node_LinkNode:
		return &ast.Link{Content: convertToASTNodes(n.LinkNode.Content), URL: n.LinkNode.Url}
	case *modelpb.Node_AutoLinkNode:
		return &ast.AutoLink{URL: n.AutoLinkNode.Url, IsRawText: n.AutoLinkNode.IsRawText}
	case *modelpb.Node_TagNode:
		return &ast.Tag{Content: n.TagNode.Content}
	case *modelpb.Node_StrikethroughNode:
		return &ast.Strikethrough{Content: n.StrikethroughNode.Content}
	case *modelpb.Node_EscapingCharacterNode:
		return &ast.EscapingCharacter{Symbol: n.EscapingCharacterNode.Symbol}
	case *modelpb.Node_MathNode:
		return &ast.Math{Content: n.MathNode.Content}
	case *modelpb.Node_HighlightNode:
		return &ast.Highlight{Content: n.HighlightNode.Content}
	case *modelpb.Node_SubscriptNode:
		return &ast.Subscript{Content: n.SubscriptNode.Content}
	case *modelpb.Node_SuperscriptNode:
		return &ast.Superscript{Content: n.SuperscriptNode.Content}
	case *modelpb.Node_ReferencedContentNode:
		return &ast.ReferencedContent{ResourceName: n.ReferencedContentNode.ResourceName, Params: n.ReferencedContentNode.Params}
	case *modelpb.Node_SpoilerNode:
		return &ast.Spoiler{Content: n.SpoilerNode.Content}
	case *modelpb.Node_HtmlElementNode:
		return &ast.HTMLElement{TagName: n.HtmlElementNode.TagName, Attributes: n.HtmlElementNode.Attributes}
	default:
		return &ast.Text{}
	}
}

func convertToASTNodes(nodes []*modelpb.Node) []ast.Node {
	rawNodes := []ast.Node{}
	for _, node := range nodes {
		rawNode := convertToASTNode(node)
		rawNodes = append(rawNodes, rawNode)
	}
	return rawNodes
}

func convertTableToASTNode(node *modelpb.TableNode) *ast.Table {
	table := &ast.Table{
		Header:    convertToASTNodes(node.Header),
		Delimiter: node.Delimiter,
	}
	for _, row := range node.Rows {
		table.Rows = append(table.Rows, convertToASTNodes(row.Cells))
	}
	return table
}

func convertListKindToASTNode(kind modelpb.ListNode_Kind) ast.ListKind {
	switch kind {
	case modelpb.ListNode_ORDERED:
		return ast.OrderedList
	case modelpb.ListNode_UNORDERED:
		return ast.UnorderedList
	case modelpb.ListNode_DESCRIPTION:
		return ast.DescrpitionList
	default:
		// Default to description list.
		return ast.DescrpitionList
	}
}
