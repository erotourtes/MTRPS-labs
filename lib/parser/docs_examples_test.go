package parser

import (
	"testing"
)

func TestDocs1(t *testing.T) {
	content := "**`_this is invalid_`**"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestDocs2(t *testing.T) {
	content := "_- this is underline"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if len(parser.GetNodes()) != 1 {
		t.Errorf("Expected 1 nodes, got %d", len(parser.GetNodes()))
	}
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestDocs3(t *testing.T) {
	content := "`_`- ok"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if len(parser.GetNodes()) != 2 {
		t.Errorf("Expected 2 nodes, got %d", len(parser.GetNodes()))
	}
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestDocs4(t *testing.T) {
	t.Run("snake_case", func(t *testing.T) {
		content := "snake_case"
		parser := MarkdownParserInit(content)
		parser.Parse()

		if len(parser.GetNodes()) != 1 {
			t.Errorf("Expected 1 nodes, got %d", len(parser.GetNodes()))
		}
	})

	t.Run("italic", func(t *testing.T) {
		content := "_italic case_"
		parser := MarkdownParserInit(content)
		parser.Parse()

		n := parser.GetNodes()[0]
		if n.Type != italic {
			t.Errorf("Expected italic, got %s", n.Type)
		}
	})
}

func TestDocs5(t *testing.T) {
	content := "_error-no ending"
	parser := MarkdownParserInit(content)
	err := parser.Parse()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
