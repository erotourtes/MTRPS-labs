package ansi

import (
	"mainmod/lib/common"
	"strings"
)

type parser = common.Parser
type node = common.Node

var mapTypeToAnsi = map[string]string{
	common.BoldT:         "\033[1m",
	common.TextT:         "",
	common.LineBreakT:    "\n",
	common.ItalicT:       "\033[3m",
	common.MonospaceT:    "\033[7m",
	common.PreformattedT: "\n\033[7m",
	"reset":              "\033[0m",
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

	return strBuilder.String()
}

func wrapIntoAnsi(node *node) string {
	return mapTypeToAnsi[node.Type] + node.Val + mapTypeToAnsi["reset"]
}
