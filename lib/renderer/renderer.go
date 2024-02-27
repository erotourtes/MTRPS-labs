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
	rendered := ""
	for _, node := range nodes {
		isStartOfLine := strings.HasSuffix(rendered, "\n")
		isEmptyText := removeRepeatedSpaces(node.Val) == ""
		if !isStartOfLine && !isEmptyText && rendered != "" {
			rendered += " "
		}
		rendered += wrapIntoTag(&node)
	}

	return "<p>\n" + rendered + "\n</p>\n"
}

func wrapIntoTag(n *Node) string {
	if n.Type == common.LineBreak {
		return "\n</p>\n<p>\n"
	} else if n.Type == common.Text {
		return removeRepeatedSpaces(n.Val)
	}
	return "<" + n.Type + ">" + removeRepeatedSpaces(n.Val) + "</" + n.Type + ">"
}

func removeRepeatedSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
