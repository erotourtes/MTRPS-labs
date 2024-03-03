package main

import (
	"fmt"
	"mainmod/lib/parser"
	"mainmod/lib/renderer"
	"mainmod/terminal"
	"os"
)

func main() {
	opt, err := terminal.GetOptions()
	terminal.ExitWithError(err)

	c := opt.GetContent()
	str, err := renderer.Render(parser.MarkdownParserInit(c))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	opt.Output(str)
}
