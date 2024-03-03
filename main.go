package main

import (
	"mainmod/lib/parser"
	"mainmod/lib/renderer/html"
	"mainmod/terminal"
)

func main() {
	opt, err := terminal.GetOptions()
	terminal.ExitIfErr(err)

	c, err := opt.GetContent()
	terminal.ExitIfErr(err)

	str, err := html.Render(parser.MarkdownParserInit(c))
	terminal.ExitIfErr(err)

	err = opt.Output(str)
	terminal.ExitIfErr(err)
}
