package parser

import (
	"testing"
)

func TestBoldSimple(t *testing.T) {
	content := "Hello **world**"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.nodes))
	}

	if parser.nodes[0].value != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.nodes[0].value)
	}

	if parser.nodes[1].value != "world" {
		t.Errorf("Expected world, got %s", parser.nodes[1].value)
	}
}

func TestBoldShouldFail(t *testing.T) {
	content := "**world"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
func TestBoldShouldFail2(t *testing.T) {
	content := "*world**"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
