package renderer

import "strings"
import "mainmod/lib/common"

type Parser = common.Parser
type Node = common.Node

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
	paragraphs := []string{}
	curParagraph := ""

	for _, node := range nodes {
		if node.Type == common.LineBreak || node.Type == common.Preformatted { // TODO
			curParagraph = "<p>\n" + curParagraph + "\n</p>\n"
			paragraphs = append(paragraphs, curParagraph)
			curParagraph = ""
		}

		isStartOfLine := strings.HasSuffix(curParagraph, "\n")
		isEmptyText := removeRepeatedSpaces(node.Val) == ""
		if !isStartOfLine && !isEmptyText && curParagraph != "" {
			curParagraph += " "
		}
		curParagraph += wrapIntoTag(&node)
	}

	if curParagraph != "" {
		paragraphs = append(paragraphs, "<p>\n"+curParagraph+"\n</p>\n")
	}

	return strings.Join(paragraphs, "\n")
}

func wrapIntoTag(n *Node) string { // TODO
	if n.Type == common.LineBreak {
		return "\n</p>\n<p>\n"
	} else if n.Type == common.Text {
		return removeRepeatedSpaces(n.Val)
	} else if n.Type == common.Preformatted {
		return "<" + n.Type + ">\n" + n.Val + "\n</" + n.Type + ">"
	}
	return "<" + n.Type + ">" + removeRepeatedSpaces(n.Val) + "</" + n.Type + ">"
}

func removeRepeatedSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}