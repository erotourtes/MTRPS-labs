package parser

import (
	"testing"
)

func TestBoldSimple(t *testing.T) {
	content := "Hello'_?' **world?**"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.GetNodes()))
	}

	n0 := parser.GetNodes()[0]
	if n0.Val != "Hello'_?' " {
		t.Errorf("Expected Hello, got %s", parser.GetNodes()[0].Val)
	}

	if n0.Type != text {
		t.Errorf("Expected text, got %s", parser.GetNodes()[0].Type)
	}

	n1 := parser.GetNodes()[1]
	if n1.Val != "world?" {
		t.Errorf("Expected world, got %s", parser.GetNodes()[1].Val)
	}

	if n1.Type != bold {
		t.Errorf("Expected bold, got %s", parser.GetNodes()[1].Type)
	}
}

func TestBoldSeveral(t *testing.T) {
	content := "Hello **world\nuniverse**"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.GetNodes()))
	}
}

func TestBoldSeveralShouldFail(t *testing.T) {
	content := "Hello **world\n\nuniverse**"
	parser := MarkdownParserInit(content)
	err := parser.parse()

	if err.line != 1 || err.col != 7 {
		t.Errorf("Expected 1, 7, got %d, %d", err.line, err.col)
	}
}

func TestBoldShouldFail(t *testing.T) {
	content := "**world"
	parser := MarkdownParserInit(content)
	err := parser.parse()
	if err.col != 1 {
		t.Errorf("Expected 1, got %d", err.col)
	}
}
func TestBoldShouldNotFail(t *testing.T) {
	content := "*world**"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestLineBreaks(t *testing.T) {
	content :=
		`Hello

text


some text`
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 5 {
		t.Errorf("Expected 5 nodes, got %d", len(parser.GetNodes()))
	}
}

func TestLineBreaks1(t *testing.T) {
	content :=
		`Hello  

text  
some text`
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 5 {
		t.Errorf("Expected 5 nodes, got %d", len(parser.GetNodes()))
	}
}

func TestItalicSimple(t *testing.T) {
	content := "Hello _world_"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.GetNodes()))
	}

	n0 := parser.GetNodes()[0]
	if n0.Val != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.GetNodes()[0].Val)
	}

	if n0.Type != text {
		t.Errorf("Expected text, got %s", parser.GetNodes()[0].Type)
	}

	n1 := parser.GetNodes()[1]
	if n1.Val != "world" {
		t.Errorf("Expected world, got %s", parser.GetNodes()[1].Val)
	}

	if n1.Type != italic {
		t.Errorf("Expected italic, got %s", parser.GetNodes()[1].Type)
	}
}

func TestItalicShouldFail(t *testing.T) {
	content := "some _world_a"
	parser := MarkdownParserInit(content)
	err := parser.parse()
	if err.col != 6 {
		t.Errorf("Expected 6, got %d", err.col)
	}
}

func TestItalicComplex(t *testing.T) {
	content := "_text_is_here_"
	parser := MarkdownParserInit(content)
	parser.Parse()
	nodes := parser.GetNodes()
	if len(nodes) > 1 {
		t.Errorf("Expected 1, got %d", len(nodes))
	}
}

func TestItalicShouldNotFail(t *testing.T) {
	content := "_ world"
	parser := MarkdownParserInit(content)
	err := parser.parse()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestItalicShouldNotFail1(t *testing.T) {
	content := "**hello_world**"
	parser := MarkdownParserInit(content)
	err := parser.parse()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestMonospaceSimple(t *testing.T) {
	content := "Hello `world`"
	parser := MarkdownParserInit(content)
	parser.Parse()

	if len(parser.GetNodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.GetNodes()))
	}

	n0 := parser.GetNodes()[0]
	if n0.Val != "Hello " {
		t.Errorf("Expected Hello, got %s", parser.GetNodes()[0].Val)
	}

	if n0.Type != text {
		t.Errorf("Expected text, got %s", parser.GetNodes()[0].Type)
	}

	n1 := parser.GetNodes()[1]
	if n1.Val != "world" {
		t.Errorf("Expected world, got %s", parser.GetNodes()[1].Val)
	}

	if n1.Type != monospace {
		t.Errorf("Expected monospace, got %s", parser.GetNodes()[1].Type)
	}
}

func TestPreformattedSimple(t *testing.T) {
	content :=
		"text\n```\n**Hello world**\n```"
	parser := MarkdownParserInit(content)
	parser.Parse()

	n0 := parser.GetNodes()[1]
	if n0.Val != "**Hello world**" {
		t.Errorf("Expected **Hello world**, got %s", parser.GetNodes()[0].Val)
	}

	if n0.Type != preformatted {
		t.Errorf("Expected preformatted, got %s", parser.GetNodes()[0].Type)
	}
}

func TestPreformattedShouldFail(t *testing.T) {
	content := "some\n```\n**Hello world**\n``"
	parser := MarkdownParserInit(content)
	err := parser.parse()
	if err.line != 2 || err.col != 1 {
		t.Errorf("Expected 2, 1, got %d, %d", err.line, err.col)
	}
}
