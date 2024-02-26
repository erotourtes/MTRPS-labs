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
}

func MarkdownParserInit(input string) *MarkdownParser {
	return &MarkdownParser{input: strings.Split(input, "\n"), nodes: []node{}}
}

func (m *MarkdownParser) Parse() error {
	for lineIdx := 0; lineIdx < len(m.input); lineIdx++ {
		line := m.input[lineIdx]
		runes := []rune(line)
		curIdx := 0
		for curIdx < len(runes) {
			if isStartOfPreformatted(runes[curIdx:]) {
				newLine, err := m.parsePreformatted(lineIdx)
				if err != nil {
					return err
				}
				lineIdx = newLine
			} else if isStartOfBold(runes[curIdx:]) {
				offset, err := m.parseStar(runes, lineIdx, curIdx)
				if err != nil {
					return err
				}
				curIdx = offset
			} else if isStartOfItalic(runes[curIdx:]) {
				offset, err := m.parseUnderscore(runes, lineIdx, curIdx)
				if err != nil {
					return err
				}
				curIdx = offset
			} else if isStartOfMonospace(runes[curIdx:]) {
				offset, err := m.parseTilda(runes, lineIdx, curIdx)
				if err != nil {
					return err
				}
				curIdx = offset
			} else {
				curIdx = m.parseText(runes, lineIdx, curIdx)
			}
		}

		if isLineBreak(line) {
			l := len(m.nodes)
			if l > 0 && m.nodes[l-1].nodeType == lineBreak {
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

func isStartOfBold(runes []rune) bool {
	return isStartOf("**", runes)
}

func isStartOfItalic(runes []rune) bool {
	return isStartOf("_", runes)
}

func isStartOfMonospace(runes []rune) bool {
	return isStartOf("`", runes)
}

func isStartOfPreformatted(runes []rune) bool {
	return strings.HasPrefix(string(runes), "```")
}

func (m *MarkdownParser) parseTilda(runes []rune, lineIdx int, startOffset int) (int, error) {
	start := startOffset + 1 // skip the first tilda
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '`' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: monospace})
			return i + 1, nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return i, fmt.Errorf("Invalid character in ` at line %d, %d ", lineIdx+1, i)
		}
	}

	return start, fmt.Errorf("no closing ` found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseStar(runes []rune, lineIdx int, startOffset int) (int, error) {
	start := startOffset + 2 // skip the first two stars
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '*' && runes[i+1] == '*' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: bold})
			return i + 2, nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return i, fmt.Errorf("Invalid character in ** at line %d, %d ", lineIdx+1, i)
		}
	}

	return start, fmt.Errorf("no closing ** found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseUnderscore(runes []rune, lineIdx int, startOffset int) (int, error) {
	start := startOffset + 1 // skip the first underscore
	for i := start; i < len(runes); i++ {
		ch := runes[i]
		if ch == '_' {
			m.nodes = append(m.nodes, node{value: string(runes[start:i]), nodeType: italic})
			return i + 1, nil
		}

		if !unicode.IsLetter(ch) && ch != ' ' {
			return i, fmt.Errorf("Invalid character in ** at line %d, %d ", lineIdx+1, i)
		}
	}

	return start, fmt.Errorf("no closing ** found at line %d, %d", lineIdx+1, startOffset)
}

func (m *MarkdownParser) parseText(runes []rune, lineIdx int, startOffset int) int {
	for i := startOffset; i < len(runes); i++ {
		ch := runes[i]
		if i != startOffset && !unicode.IsLetter(ch) && ch != ' ' {
			m.nodes = append(m.nodes, node{value: string(runes[startOffset:i]), nodeType: text})
			return i
		}
	}

	m.nodes = append(m.nodes, node{value: string(runes[startOffset:]), nodeType: text})
	return len(runes)
}

func (m *MarkdownParser) parsePreformatted(lineIdx int) (int, error) {
	line := m.input[lineIdx]
	if len(line) < 3 {
		return 0, fmt.Errorf("invalid preformatted block at line %d", lineIdx+1)
	}

	for i := lineIdx + 1; i < len(m.input); i++ {
		line = m.input[i]
		if strings.HasPrefix(line, "```") {
			if len(line) > 3 {
				return 0, fmt.Errorf("invalid preformatted block at line %d", i+1)
			}
			m.nodes = append(
				m.nodes,
				node{value: strings.Join(m.input[lineIdx+1:i], "\n"), nodeType: preformatted})
			return i, nil
		}
	}

	return 0, fmt.Errorf("no closing ``` found at line %d", lineIdx+1)
}
