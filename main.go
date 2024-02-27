package main

import (
	"fmt"
	"mainmod/parser"
	"mainmod/renderer"
)

func main() {
	content := "Hello **world**"
	str, err := renderer.Render(content, parser.MarkdownParserInit(content))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(str)
}
