package main

import (
	"fmt"
	"mainmod/parser"
)

func main() {
	// const path = "./README.md"
	// content, _ := getMDContentFrom(path)
	content := "Hello **world**"
	parser := parser.MarkdownParserInit(content)
	parser.Parse()

	fmt.Println(content)
}
