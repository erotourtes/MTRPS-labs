package ansi

import (
	"mainmod/lib/common"
	p "mainmod/lib/parser"
	"testing"
)

func _w(type_ string, val string) string {
	return mapTypeToAnsi[type_] + val + mapTypeToAnsi["reset"]
}

func TestRender(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple",
			input:    `Hello **world**!`,
			expected: _w(common.TextT, "Hello ") + _w(common.BoldT, "world") + _w(common.TextT, "!"),
		},
		{
			name:     "Different types (bold and italic)",
			input:    `**Hello** *world  ` + "`A`",
			expected: _w(common.BoldT, "Hello") + _w(common.TextT, " *world  ") + _w(common.MonospaceT, "A"),
		},
		{
			name:  "Preformatted",
			input: "```\n**damn**\n```\n```\nHello\nworld\n```",
			expected: _w(common.PreformattedT, "**damn**") +
				_w(common.PreformattedT, "Hello\nworld"),
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
