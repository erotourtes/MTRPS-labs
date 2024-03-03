package ansi

import (
	"fmt"
	"mainmod/lib/common"
	"strings"
)

type parser = common.Parser
type node = common.Node

var mapTypeToAnsi = map[string]string{
	common.BoldT:         "\033[1m%s\033[0m",
	common.TextT:         "%s",
	common.LineBreakT:    "%s\n",
	common.ItalicT:       "\033[3m%s\033[0m",
	common.MonospaceT:    "\033[7m%s\033[0m",
	common.PreformattedT: "\n\033[7m%s\033[0m\n",
}

func Render(parser parser) (string, error) {
	err := parser.Parse()
	if err != nil {
		return "", err
	}

	nodes := parser.GetNodes()
	rendered := renderNodes(nodes)

	return rendered, nil
}

func renderNodes(nodes []node) string {
	strBuilder := new(strings.Builder)

	for _, node := range nodes {
		strBuilder.WriteString(wrapIntoAnsi(&node))
	}

	return strings.TrimSpace(strBuilder.String())
}

func wrapIntoAnsi(node *node) string {
	if node.Type == common.PreformattedT || node.Type == common.MonospaceT {
		return fmt.Sprintf(mapTypeToAnsi[node.Type], node.Val)
	}
	return fmt.Sprintf(mapTypeToAnsi[node.Type], common.RemoveRepeatedSpaces(node.Val))
}
