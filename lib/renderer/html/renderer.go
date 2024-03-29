package html

import (
	"mainmod/lib/common"
	"strings"
)

type Parser = common.Parser
type Node = common.Node

var mapTypeToTag = map[string]string{
	common.BoldT:         "b",
	common.TextT:         "text",
	common.LineBreakT:    "lineBreak",
	common.ItalicT:       "i",
	common.MonospaceT:    "tt",
	common.PreformattedT: "pre",
}

func Render(parser Parser) (string, error) {
	err := parser.Parse()
	if err != nil {
		return "", err
	}

	nodes := parser.GetNodes()
	rendered := renderNodes(nodes)

	return rendered, nil
}

func renderNodes(nodes []Node) string {
	var paragraphs []string
	curParagraph := ""

	for _, node := range nodes {
		if node.Type == common.LineBreakT {
			if isOnlySpaces := strings.TrimSpace(curParagraph) == ""; !isOnlySpaces {
				curParagraph = wrapIntoParagraph(curParagraph)
				paragraphs = append(paragraphs, curParagraph)
			}
			curParagraph = ""
		}

		curParagraph += wrapIntoTag(&node)
	}

	if curParagraph != "" {
		paragraphs = append(paragraphs, wrapIntoParagraph(curParagraph))
	}

	return strings.Join(paragraphs, "\n")
}

func wrapIntoParagraph(s string) string {
	return "<p>\n" + strings.TrimSpace(s) + "\n</p>\n"
}

func wrapIntoTag(n *Node) string { // TODO
	if n.Type == common.LineBreakT {
		return ""
	} else if n.Type == common.TextT {
		return common.RemoveRepeatedSpaces(n.Val)
	} else if n.Type == common.PreformattedT {
		return "\n<" + symb(n.Type) + ">\n" + n.Val + "\n</" + symb(n.Type) + ">\n"
	}
	return "<" + symb(n.Type) + ">" + common.RemoveRepeatedSpaces(n.Val) + "</" + symb(n.Type) + ">"
}

func symb(typ string) string {
	return mapTypeToTag[typ]
}
