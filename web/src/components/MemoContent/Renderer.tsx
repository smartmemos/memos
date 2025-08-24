import {
  AutoLinkNode,
  BlockquoteNode,
  BoldItalicNode,
  BoldNode,
  CodeBlockNode,
  CodeNode,
  EmbeddedContentNode,
  EscapingCharacterNode,
  HeadingNode,
  HighlightNode,
  HorizontalRuleNode,
  HTMLElementNode,
  ImageNode,
  ItalicNode,
  LinkNode,
  ListNode,
  MathBlockNode,
  MathNode,
  Node,
  NodeType,
  OrderedListItemNode,
  ParagraphNode,
  ReferencedContentNode,
  SpoilerNode,
  StrikethroughNode,
  SubscriptNode,
  SuperscriptNode,
  TableNode,
  TagNode,
  TaskListItemNode,
  TextNode,
  UnorderedListItemNode,
} from "@/types/proto2/model/markdown_pb";
import Blockquote from "./Blockquote";
import Bold from "./Bold";
import BoldItalic from "./BoldItalic";
import Code from "./Code";
import CodeBlock from "./CodeBlock";
import EmbeddedContent from "./EmbeddedContent";
import EscapingCharacter from "./EscapingCharacter";
import HTMLElement from "./HTMLElement";
import Heading from "./Heading";
import Highlight from "./Highlight";
import HorizontalRule from "./HorizontalRule";
import Image from "./Image";
import Italic from "./Italic";
import LineBreak from "./LineBreak";
import Link from "./Link";
import List from "./List";
import Math from "./Math";
import OrderedListItem from "./OrderedListItem";
import Paragraph from "./Paragraph";
import ReferencedContent from "./ReferencedContent";
import Spoiler from "./Spoiler";
import Strikethrough from "./Strikethrough";
import Subscript from "./Subscript";
import Superscript from "./Superscript";
import Table from "./Table";
import Tag from "./Tag";
import TaskListItem from "./TaskListItem";
import Text from "./Text";
import UnorderedListItem from "./UnorderedListItem";

interface Props {
  index: string;
  node: Node;
}

const Renderer: React.FC<Props> = ({ index, node }: Props) => {
  switch (node.type) {
    case NodeType.LINE_BREAK:
      return <LineBreak />;
    case NodeType.PARAGRAPH:
      return <Paragraph index={index} {...(node.node.value as ParagraphNode)} />;
    case NodeType.CODE_BLOCK:
      return <CodeBlock index={index} {...(node.node.value as CodeBlockNode)} />;
    case NodeType.HEADING:
      return <Heading index={index} {...(node.node.value as HeadingNode)} />;
    case NodeType.HORIZONTAL_RULE:
      return <HorizontalRule index={index} {...(node.node.value as HorizontalRuleNode)} />;
    case NodeType.BLOCKQUOTE:
      return <Blockquote index={index} {...(node.node.value as BlockquoteNode)} />;
    case NodeType.LIST:
      return <List index={index} {...(node.node.value as ListNode)} />;
    case NodeType.ORDERED_LIST_ITEM:
      return <OrderedListItem index={index} {...(node.node.value as OrderedListItemNode)} />;
    case NodeType.UNORDERED_LIST_ITEM:
      return <UnorderedListItem {...(node.node.value as UnorderedListItemNode)} />;
    case NodeType.TASK_LIST_ITEM:
      return <TaskListItem index={index} node={node} {...(node.node.value as TaskListItemNode)} />;
    case NodeType.MATH_BLOCK:
      return <Math {...(node.node.value as MathBlockNode)} block={true} />;
    case NodeType.TABLE:
      return <Table index={index} {...(node.node.value as TableNode)} />;
    case NodeType.EMBEDDED_CONTENT:
      return <EmbeddedContent {...(node.node.value as EmbeddedContentNode)} />;
    case NodeType.TEXT:
      return <Text {...(node.node.value as TextNode)} />;
    case NodeType.BOLD:
      return <Bold {...(node.node.value as BoldNode)} />;
    case NodeType.ITALIC:
      return <Italic {...(node.node.value as ItalicNode)} />;
    case NodeType.BOLD_ITALIC:
      return <BoldItalic {...(node.node.value as BoldItalicNode)} />;
    case NodeType.CODE:
      return <Code {...(node.node.value as CodeNode)} />;
    case NodeType.IMAGE:
      return <Image {...(node.node.value as ImageNode)} />;
    case NodeType.LINK:
      return <Link {...(node.node.value as LinkNode)} />;
    case NodeType.AUTO_LINK:
      return <Link {...(node.node.value as AutoLinkNode)} />;
    case NodeType.TAG:
      return <Tag {...(node.node.value as TagNode)} />;
    case NodeType.STRIKETHROUGH:
      return <Strikethrough {...(node.node.value as StrikethroughNode)} />;
    case NodeType.MATH:
      return <Math {...(node.node.value as MathNode)} />;
    case NodeType.HIGHLIGHT:
      return <Highlight {...(node.node.value as HighlightNode)} />;
    case NodeType.ESCAPING_CHARACTER:
      return <EscapingCharacter {...(node.node.value as EscapingCharacterNode)} />;
    case NodeType.SUBSCRIPT:
      return <Subscript {...(node.node.value as SubscriptNode)} />;
    case NodeType.SUPERSCRIPT:
      return <Superscript {...(node.node.value as SuperscriptNode)} />;
    case NodeType.REFERENCED_CONTENT:
      return <ReferencedContent {...(node.node.value as ReferencedContentNode)} />;
    case NodeType.SPOILER:
      return <Spoiler {...(node.node.value as SpoilerNode)} />;
    case NodeType.HTML_ELEMENT:
      return <HTMLElement {...(node.node.value as HTMLElementNode)} />;
    default:
      return null;
  }
};

export default Renderer;
