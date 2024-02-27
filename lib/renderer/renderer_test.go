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
`,
			expected: `<p>
Hello <b>world</b>
</p>
`},
		{
			name: "Different types (bold and italic)",
			input: `

**Hello** *world  
` + "`A`" + `

`,
			expected: `<p>
<b>Hello</b> *world
</p>
<p>
<tt>A</tt>
</p>
`,
		},
		//		{
		//			name: "Different types (bold, italic, monospace, preformatted)",
		//			input: `
		//**Hello** *world* ` + "`A`" + "\n```" + `
		//**H** _A <>
		//A_` + "\n```" + `
		//`,
		//			expected: `<p>
		//<b>Hello</b> <i>world</i> <tt>A</tt>
		//</p>
		//<pre>
		//**H** _A <>
		//A_
		//</pre>
		//`,
		//		},
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
