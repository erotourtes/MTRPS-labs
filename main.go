package main

import (
	"fmt"
	"mainmod/lib/parser"
	"mainmod/lib/renderer"
	"mainmod/terminal"
	"os"
)

func main() {
	opt := terminal.GetOptions()
	c := opt.GetContent()
	str, err := renderer.Render(parser.MarkdownParserInit(c))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	opt.Output(str)
}
