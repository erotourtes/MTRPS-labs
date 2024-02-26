package parser

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	bold         = "b"
	text         = "text"
	lineBreak    = "lineBreak"
	italic       = "i"
	monospace    = "tt"
	preformatted = "pre"
)

type node struct {
	value    string
	nodeType string
}

type MarkdownParser struct {
	input []string
	nodes []node

	pos struct {
		line int
		col  int
	}
}

func (m *MarkdownParser) curLine() string {
	return m.input[m.pos.line]
}

func (m *MarkdownParser) curLineRunes() []rune {
	return []rune(m.curLine())
}

func (m *MarkdownParser) lastNodeType() string {
	l := len(m.nodes)
	if l == 0 {
		return ""
	}
	return m.nodes[l-1].nodeType
}

func MarkdownParserInit(input string) *MarkdownParser {
	return &MarkdownParser{input: strings.Split(input, "\n"), nodes: []node{}}
}

func (m *MarkdownParser) Parse() error {
	for ; m.pos.line < len(m.input); m.pos.line++ {
		line := m.curLine()
		runes := m.curLineRunes()
		for m.pos.col < len(runes) {
			var err error
			if m.isStartOfPreformatted() {
				err = m.parsePreformatted()
			} else if m.isStartOfBold() {
				err = m.parseStar(runes)
			} else if m.isStartOfItalic() {
				err = m.parseUnderscore(runes)
			} else if m.isStartOfMonospace() {
				err = m.parseTilda(runes)
			} else {
				m.parseText(runes)
			}
			if err != nil {
				return err
			}
		}

		m.pos.col = 0

		if isLineBreak(line) {
			if m.lastNodeType() == lineBreak {
				continue
			}
			m.nodes = append(m.nodes, node{value: "", nodeType: lineBreak})
		}
	}

	return nil
}

func isLineBreak(line string) bool {
	if strings.HasSuffix(line, "  ") {
		return true
	}
	return len(strings.TrimSpace(line)) == 0
}

/*
Helper function to check if the given string is the start of the md syntax
*/
func isStartOf(str string, runes []rune) bool {
	if len(runes) < len(str) {
		return false
	}

	runesStr := string(runes)
	return strings.HasPrefix(runesStr, str) && (len(runes) > len(str) && unicode.IsLetter(runes[len(str)]))
}

func (m *MarkdownParser) isStartOfBold() bool {
	runes := m.curLineRunes()[m.pos.col:]
	return isStartOf("**", runes)
}

func (m *MarkdownParser) isStartOfItalic() bool {
	runes := m.curLineRunes()[m.pos.col:]
	return isStartOf("_", runes)
}

func (m *MarkdownParser) isStartOfMonospace() bool {
	runes := m.curLineRunes()[m.pos.col:]
	return isStartOf("`", runes)
}

func (m *MarkdownParser) isStartOfPreformatted() bool {
	runes := m.curLineRunes()[m.pos.col:]
	return strings.HasPrefix(string(runes), "```")
}

func (m *MarkdownParser) parseTilda(runes []rune) error {
	lineIdx := m.pos.line
	startOffset := m.pos.col

	start := startOffset + 1 // skip the first tilda
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '`' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: monospace})
			m.pos.col = i + 1
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return fmt.Errorf("Invalid character in ` at line %d, %d ", lineIdx+1, i)
		}
	}

	return fmt.Errorf("no closing ` found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseStar(runes []rune) error {
	lineIdx := m.pos.line
	startOffset := m.pos.col

	start := startOffset + 2 // skip the first two stars
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '*' && runes[i+1] == '*' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: bold})
			m.pos.col = i + 2
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return fmt.Errorf("Invalid character in ** at line %d, %d ", lineIdx+1, i)
		}
	}

	return fmt.Errorf("no closing ** found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseUnderscore(runes []rune) error {
	lineIdx := m.pos.line
	startOffset := m.pos.col

	start := startOffset + 1 // skip the first underscore
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '_' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: italic})
			m.pos.col = i + 1
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return fmt.Errorf("Invalid character in ** at line %d, %d ", lineIdx+1, i)
		}
	}

	return fmt.Errorf("no closing ** found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseText(runes []rune) {
	startOffset := m.pos.col
	for i := startOffset; i < len(runes); i++ {
		ch := runes[i]
		if i != startOffset && !unicode.IsLetter(ch) && ch != ' ' {
			m.nodes = append(m.nodes, node{value: string(runes[startOffset:i]), nodeType: text})
			m.pos.col = i
			return
		}
	}

	m.nodes = append(m.nodes, node{value: string(runes[startOffset:]), nodeType: text})
	m.pos.col = len(runes)
}

func (m *MarkdownParser) parsePreformatted() error {
	lineIdx := m.pos.line
	line := m.input[lineIdx]
	if len(line) < 3 {
		return fmt.Errorf("invalid preformatted block at line %d", lineIdx+1)
	}

	for i := lineIdx + 1; i < len(m.input); i++ {
		line = m.input[i]
		if strings.HasPrefix(line, "```") {
			if len(line) > 3 {
				return fmt.Errorf("invalid preformatted block at line %d", i+1)
			}
			m.nodes = append(
				m.nodes,
				node{value: strings.Join(m.input[lineIdx+1:i], "\n"), nodeType: preformatted})
			m.pos.line = i
			return nil
		}
	}

	return fmt.Errorf("no closing ``` found at line %d", lineIdx+1)
}
