package html

import (
	"regexp"
	"strings"
)
import "mainmod/lib/common"

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
		if node.Type == common.LineBreakT || node.Type == common.PreformattedT { // TODO
			isOnlySpaces := strings.TrimSpace(curParagraph) == ""
			if !isOnlySpaces {
				curParagraph = wrapIntoParagraph(curParagraph)
				paragraphs = append(paragraphs, curParagraph)
			}
			curParagraph = ""

			if node.Type == common.PreformattedT {
				paragraphs = append(paragraphs, wrapIntoParagraph(wrapIntoTag(&node)))
				continue
			}
		}

		curParagraph += wrapIntoTag(&node)
	}

	if curParagraph != "" {
		paragraphs = append(paragraphs, wrapIntoParagraph(curParagraph))
	}

	return strings.Join(paragraphs, "\n")
}

func wrapIntoParagraph(s string) string {
	return "<p>\n" + s + "\n</p>\n"
}

func wrapIntoTag(n *Node) string { // TODO
	if n.Type == common.LineBreakT {
		return ""
	} else if n.Type == common.TextT {
		return removeRepeatedSpaces(n.Val)
	} else if n.Type == common.PreformattedT {
		return "<" + symb(n.Type) + ">\n" + n.Val + "\n</" + symb(n.Type) + ">"
	}
	return "<" + symb(n.Type) + ">" + removeRepeatedSpaces(n.Val) + "</" + symb(n.Type) + ">"
}

func removeRepeatedSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	output := re.ReplaceAllString(s, " ")
	return output
}

func symb(typ string) string {
	return mapTypeToTag[typ]
}
