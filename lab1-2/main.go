package main

import (
	"mainmod/lib/parser"
	"mainmod/terminal"
)

func main() {
	opt, err := terminal.GetOptions()
	terminal.ExitIfErr(err)

	c, err := opt.GetContent()
	terminal.ExitIfErr(err)

	str, err := opt.RenderWith(parser.MarkdownParserInit(c))
	terminal.ExitIfErr(err)

	err = opt.Output(str)
	terminal.ExitIfErr(err)
}
