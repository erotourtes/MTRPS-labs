package renderer

import (
	"mainmod/lib/parser"
	"testing"
)

func TestRender(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Simple",
			input: `
Hello **world**
`, expected: `
<p>
Hello <b>world</b>
</p>
`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			str, err := Render(parser.MarkdownParserInit(tc.input))
			if err != nil {
				t.Error(err)
			}
			if str != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, str)
			}
		})
	}
}
