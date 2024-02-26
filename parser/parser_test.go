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

	n0 := parser.nodes[0]
	if n0.value != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.nodes[0].value)
	}

	if n0.nodeType != text {
		t.Errorf("Expected text, got %s", parser.nodes[0].nodeType)
	}

	n1 := parser.nodes[1]
	if n1.value != "world" {
		t.Errorf("Expected world, got %s", parser.nodes[1].value)
	}

	if n1.nodeType != bold {
		t.Errorf("Expected bold, got %s", parser.nodes[1].nodeType)
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

func TestLineBreaks(t *testing.T) {
	content :=
		`Hello

text


some text`
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.nodes) != 5 {
		t.Errorf("Expected 5 nodes, got %d", len(parser.nodes))
	}
}

func TestLineBreaks1(t *testing.T) {
	content :=
		`Hello  

text  
some text`
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.nodes) != 5 {
		t.Errorf("Expected 5 nodes, got %d", len(parser.nodes))
	}
}

func TestItalicSimple(t *testing.T) {
	content := "Hello _world_"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.nodes))
	}

	n0 := parser.nodes[0]
	if n0.value != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.nodes[0].value)
	}

	if n0.nodeType != text {
		t.Errorf("Expected text, got %s", parser.nodes[0].nodeType)
	}

	n1 := parser.nodes[1]
	if n1.value != "world" {
		t.Errorf("Expected world, got %s", parser.nodes[1].value)
	}

	if n1.nodeType != italic {
		t.Errorf("Expected italic, got %s", parser.nodes[1].nodeType)
	}
}

func TestItalicShouldFail(t *testing.T) {
	content := "_world"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestItalicShouldNotFail(t *testing.T) {
	content := "_ world"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestMonospaceSimple(t *testing.T) {
	content := "Hello `world`"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.nodes) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.nodes))
	}

	n0 := parser.nodes[0]
	if n0.value != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.nodes[0].value)
	}

	if n0.nodeType != text {
		t.Errorf("Expected text, got %s", parser.nodes[0].nodeType)
	}

	n1 := parser.nodes[1]
	if n1.value != "world" {
		t.Errorf("Expected world, got %s", parser.nodes[1].value)
	}

	if n1.nodeType != monospace {
		t.Errorf("Expected monospace, got %s", parser.nodes[1].nodeType)
	}
}
