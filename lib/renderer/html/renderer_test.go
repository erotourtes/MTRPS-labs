package html

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
		{
			name: "Different types (bold, italic, monospace, preformatted)",
			/*

				*Hello** *world* `A`
				```
				**H** _A <>
				A_
				```

			*/
			input: `
**Hello** *world*` + "`A`" + "\n```" + `
**H** _A <>
A_` + "\n```\n",
			expected: `<p>
<b>Hello</b> *world*<tt>A</tt>
</p>

<p>
<pre>
**H** _A <>
A_
</pre>
</p>
`,
		},
		{
			name: "Preformatted",
			input: `
` + "```" + `
a
` + "```" + `
b
` + "```" + `
c
` + "```" + `
`,
			expected: `<p>
<pre>
a
</pre>
</p>

<p>
b
</p>

<p>
<pre>
c
</pre>
</p>
`,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			str, err := Render(parser.MarkdownParserInit(tc.input))
			if err != nil {
				t.Fatal(err)
			}
			if str != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, str)
			}
		})
	}
}
