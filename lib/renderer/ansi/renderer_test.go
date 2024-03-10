package ansi

import (
	p "mainmod/lib/parser"
	"testing"
)

func TestRender(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple",
			input:    `Hello **world**!`,
			expected: "Hello \033[1mworld\033[0m!",
		},
		{
			name:     "Different types (bold and italic)",
			input:    `**Hello** *world  ` + "`A`",
			expected: "\033[1mHello\033[0m *world \033[7mA\033[0m",
		},
		{
			name:     "Preformatted",
			input:    "```\n**damn**\n```\n```\nHello\nworld\n```",
			expected: "\033[7m**damn**\033[0m\n\n\033[7mHello\nworld\033[0m",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			parser := p.MarkdownParserInit(tc.input)
			rendered, err := Render(parser)
			if err != nil {
				t.Errorf("Error: %s", err)
			}
			if rendered != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, rendered)
			}
		})
	}
}
