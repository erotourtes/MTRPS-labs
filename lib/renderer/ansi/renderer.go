package ansi

import "mainmod/lib/common"

type Parser = common.Parser
type Node = common.Node

var mapTypeToAnsi = map[string]string{
	common.BoldT:         "\033[1m",
	common.TextT:         "",
	common.LineBreakT:    "\n",
	common.ItalicT:       "\033[3m",
	common.MonospaceT:    "\033[27m",
	common.PreformattedT: "\033[27m",
	"reset":              "\033[0m",
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
	return ""
}
