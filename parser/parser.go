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

type ParserError struct {
	line int
	col  int
	msg  string
}

func (p *ParserError) Error() string {
	return fmt.Sprintf("Error at line %d, col %d: %s", p.line, p.col, p.msg)
}

type MarkdownParser struct {
	input []string
	nodes []node

	pos struct {
		line int
		col  int
	}
}

func (m *MarkdownParser) curStrLine() string {
	return m.input[m.pos.line]
}

func (m *MarkdownParser) curLineRunes() []rune {
	return []rune(m.curStrLine())
}

func (m *MarkdownParser) lastNodeType() string {
	l := len(m.nodes)
	if l == 0 {
		return ""
	}
	return m.nodes[l-1].nodeType
}

func (m *MarkdownParser) getLine() int {
	return m.pos.line
}

func (m *MarkdownParser) getCol() int {
	return m.pos.col
}

func (m *MarkdownParser) setLine(line int) {
	m.pos.line = line
}

func (m *MarkdownParser) setCol(col int) {
	m.pos.col = col
}

func (m *MarkdownParser) incrementLine() {
	m.pos.line++
}

func MarkdownParserInit(input string) *MarkdownParser {
	return &MarkdownParser{input: strings.Split(input, "\n"), nodes: []node{}}
}

func (m *MarkdownParser) error(msg string) *ParserError {
	return &ParserError{line: m.getLine() + 1, col: m.getCol() + 1, msg: msg}
}

func (m *MarkdownParser) Parse() *ParserError {
	for ; m.getLine() < len(m.input); m.incrementLine() {
		runes := m.curLineRunes()
		for m.getCol() < len(runes) {
			var err *ParserError
			switch true {
			case m.isStartOfPreformatted():
				err = m.parsePreformatted()
			case m.isStartOfBold():
				err = m.parseStar()
			case m.isStartOfItalic():
				err = m.parseUnderscore()
			case m.isStartOfMonospace():
				err = m.parseTilda()
			default:
				m.parseText()
			}
			if err != nil {
				return err
			}
		}

		m.setCol(0)

		if m.isLineBreak() {
			if m.lastNodeType() == lineBreak {
				continue
			}
			m.nodes = append(m.nodes, node{value: "", nodeType: lineBreak})
		}
	}

	return nil
}

func (m *MarkdownParser) isLineBreak() bool {
	line := m.curStrLine()
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
	runes := m.curLineRunes()[m.getCol():]
	return isStartOf("**", runes)
}

func (m *MarkdownParser) isStartOfItalic() bool {
	runes := m.curLineRunes()[m.getCol():]
	return isStartOf("_", runes)
}

func (m *MarkdownParser) isStartOfMonospace() bool {
	runes := m.curLineRunes()[m.getCol():]
	return isStartOf("`", runes)
}

func (m *MarkdownParser) isStartOfPreformatted() bool {
	runes := m.curLineRunes()[m.getCol():]
	return strings.HasPrefix(string(runes), "```")
}

func (m *MarkdownParser) parseTilda() *ParserError {
	runes := m.curLineRunes()
	start := m.getCol() + 1 // skip the first tilda
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '`' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: monospace})
			m.setCol(i + 1)
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return m.error("Invalid character in `")
		}
	}

	return m.error("No closing ` found")
}

func (m *MarkdownParser) parseStar() *ParserError {
	runes := m.curLineRunes()
	start := m.getCol() + 2 // skip the first two stars
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '*' && runes[i+1] == '*' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: bold})
			m.setCol(i + 2)
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return m.error("Invalid character in **")
		}
	}

	return m.error("No closing ** found")
}

func (m *MarkdownParser) parseUnderscore() *ParserError {
	runes := m.curLineRunes()
	start := m.getCol() + 1 // skip the first underscore
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '_' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: italic})
			m.setCol(i + 1)
			return nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return m.error("Invalid character in _")
		}
	}

	return m.error("No closing _ found")
}

func (m *MarkdownParser) parseText() {
	runes := m.curLineRunes()
	startOffset := m.getCol()
	for i := startOffset; i < len(runes); i++ {
		ch := runes[i]
		if i != startOffset && !unicode.IsLetter(ch) && ch != ' ' {
			m.nodes = append(m.nodes, node{value: string(runes[startOffset:i]), nodeType: text})
			m.setCol(i)
			return
		}
	}

	m.nodes = append(m.nodes, node{value: string(runes[startOffset:]), nodeType: text})
	m.setCol(len(runes))
}

func (m *MarkdownParser) parsePreformatted() *ParserError {
	lineIdx := m.getLine()
	line := m.input[lineIdx]
	if len(line) < 3 {
		return m.error("Invalid preformatted block")
	}

	for i := m.getLine() + 1; i < len(m.input); i++ {
		line = m.input[i]
		if strings.HasPrefix(line, "```") {
			if len(line) > 3 {
				return m.error("Invalid preformatted block")
			}
			m.nodes = append(
				m.nodes,
				node{value: strings.Join(m.input[lineIdx+1:i], "\n"), nodeType: preformatted})
			m.setLine(i)
			return nil
		}
	}

	return m.error("No closing ``` found")
}
