package parser

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	bold = "b"
	text = "text"
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
			if isStartOfBold(runes[curIdx:]) {
				offset, err := m.parseStar(runes, lineIdx, curIdx)
				if err != nil {
					return err
				}
				curIdx = offset
			} else {
				curIdx = m.parseText(runes, lineIdx, curIdx)
			}
		}
	}

	return nil
}

func isStartOfBold(runes []rune) bool {
	if len(runes) < 2 {
		return false
	}

	return runes[0] == '*' && runes[1] == '*' && (len(runes) == 2 || unicode.IsLetter(runes[2]))
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
