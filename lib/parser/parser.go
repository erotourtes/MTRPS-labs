package parser

import (
	"fmt"
	"mainmod/lib/common"
	"strings"
	"unicode"
)

type Node = common.Node
type Pos = common.Pos

const lineBreak = common.LineBreak
const bold = common.Bold
const text = common.Text
const italic = common.Italic
const monospace = common.Monospace
const preformatted = common.Preformatted

var mapHandlers = map[string]func(m *MarkdownParser, node *Node) *ParserError{
	"```": (*MarkdownParser).parsePreformatted,
	"_":   (*MarkdownParser).parseUnderscore,
	"**":  (*MarkdownParser).parseStar,
	"`":   (*MarkdownParser).parseTilda,
}

var mapTypes = map[string]string{
	"```": preformatted,
	"**":  bold,
	"_":   italic,
	"`":   monospace,
}

var mapTypesRev = map[string]string{
	preformatted: "```",
	bold:         "**",
	italic:       "_",
	monospace:    "`",
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
	root  *Node
}

func (m *MarkdownParser) GetNodes() []Node {
	return m.root.Children
}

func (m *MarkdownParser) Parse() error {
	// TODO: Wut? can't directly return the error?
	err := m.parse()
	if err != nil {
		return err
	}
	return nil
}

func (m *MarkdownParser) curStrLine() string {
	return m.input[m.root.Pos.Line]
}

func (m *MarkdownParser) curLineRunes() []rune {
	return []rune(m.curStrLine())
}

func (m *MarkdownParser) getLine() int {
	return m.root.Pos.Line
}

func (m *MarkdownParser) getCol() int {
	return m.root.Pos.Col
}

func (m *MarkdownParser) setLine(line int) {
	m.root.Pos.Line = line
}

func (m *MarkdownParser) setCol(col int) {
	m.root.Pos.Col = col
}

func (m *MarkdownParser) incrementLine() {
	m.root.Pos.Line++
}

func (m *MarkdownParser) incrementColBy(val int) {
	m.root.Pos.Col += val
}

func (m *MarkdownParser) incrementCol() {
	m.incrementColBy(1)
}

func MarkdownParserInit(input string) *MarkdownParser {
	return &MarkdownParser{input: strings.Split(input, "\n"), root: &Node{Type: "root", Children: []Node{}, Pos: &Pos{Line: 0, Col: 0, EndLine: 0, EndCol: 0}}}
}

func errorFor(node *Node, msg string) *ParserError {
	typLen := len(mapTypesRev[node.Type])
	return &ParserError{line: node.Pos.Line + 1, col: node.Pos.Col + 1 - typLen, msg: msg}
}

func (m *MarkdownParser) error(msg string) *ParserError {
	return &ParserError{line: m.getLine() + 1, col: m.getCol() + 1, msg: msg}
}

/*
Works without taking to account inner nodes (because there will be none as defined in the task description)
*/
func (m *MarkdownParser) lastNodeType() string {
	if len(m.root.Children) == 0 {
		return ""
	}
	return m.root.Children[len(m.root.Children)-1].Type
}

func (m *MarkdownParser) newNode(typ string) *Node {
	return &Node{Type: typ, Children: []Node{}, Pos: &Pos{
		Line:    m.getLine(),
		Col:     m.getCol(),
		EndLine: m.getLine(),
		EndCol:  m.getCol(),
	}}
}

func (m *MarkdownParser) parse() *ParserError {
	node := m.root
	for ; m.getLine() < len(m.input); m.incrementLine() {
		for m.setCol(0); m.getCol() < len(m.curLineRunes()); {
			for typ, handler := range mapHandlers {
				if m.isStartOf(typ) {
					err := handler(m, node)
					if err != nil {
						return err
					}
					break
				}
			}
			// none of the special types matched
			if m.getCol() < len(m.curLineRunes()) {
				m.parseText(node)
			}
		}

		if m.isLineBreak() {
			if m.lastNodeType() == lineBreak {
				continue
			}
			newNode := m.newNode(lineBreak)
			m.closeNode(newNode)
			m.root.Children = append(m.root.Children, *newNode)
		}
	}

	return nil
}

func (m *MarkdownParser) getValOf(pos *Pos) string {
	if pos.Line == pos.EndLine {
		return m.input[pos.Line][pos.Col:pos.EndCol]
	}
	return m.input[pos.Line][pos.Col:] + "\n" + strings.Join(m.input[pos.Line+1:pos.EndLine], "\n") + m.input[pos.EndLine][:pos.EndCol]
}

func (m *MarkdownParser) isLineBreak() bool {
	line := m.curStrLine()
	return strings.HasSuffix(line, "  ") || len(strings.TrimSpace(line)) == 0
}

func (m *MarkdownParser) isStartOf(typ string) bool {
	runes := m.curLineRunes()[m.getCol():]
	// TODO: define start of separately
	if typ == "```" {
		return strings.HasPrefix(string(runes), "```")
	} else if typ == "_" {
		// this shouldn't be start of italic **hello_world**
		isPrevLetter := m.getCol() > 0 && unicode.IsLetter(m.curLineRunes()[m.getCol()-1])
		return !isPrevLetter && strings.HasPrefix(string(runes), typ) &&
			(len(runes) > len(typ) && unicode.IsLetter(runes[len(typ)]))
	} else if typ == "`" {
		return strings.HasPrefix(string(runes), typ)
	}

	if len(runes) < len(typ) {
		return false
	}

	return strings.HasPrefix(string(runes), typ) &&
		(len(runes) > len(typ) && unicode.IsLetter(runes[len(typ)]))
}

func (m *MarkdownParser) isStartOfAnotherType() bool {
	for typ := range mapTypes {
		if m.isStartOf(typ) {
			return true
		}
	}

	return false
}

func (m *MarkdownParser) isEndOf(typ string) bool {
	runes := m.curLineRunes()[m.getCol():]
	if len(runes) < len(typ) {
		return false
	}

	return strings.HasPrefix(string(runes), typ)
}

func (m *MarkdownParser) closeNode(node *Node) {
	node.Pos.EndLine = m.getLine()
	node.Pos.EndCol = m.getCol()
	node.IsClosed = true
	node.Val = m.getValOf(node.Pos)
}

func (m *MarkdownParser) parseDefault(typ string, parent *Node, treatEndSymbolBeforeLetter bool) *ParserError {
	m.incrementColBy(len(typ)) // skip the starting symbols
	typName := mapTypes[typ]
	newNode := m.newNode(typName)
	for ; m.getLine() < len(m.input); m.incrementLine() {
		if isEmptyLine := len(strings.TrimSpace(m.input[m.getLine()])) == 0; isEmptyLine {
			return errorFor(newNode, fmt.Sprintf("No closing %s found", typ))
		}
		if m.getLine() != newNode.Pos.Line {
			m.setCol(0) // reset the column if we are on a new line
		}
		runes := m.curLineRunes()
		for ; m.getCol() < len(runes); m.incrementCol() {
			if m.isEndOf(typ) {
				// with treatEndSymbolBeforeLetter: false this is valid _hello_world_
				if !treatEndSymbolBeforeLetter && len(runes) > m.getCol()+1 && unicode.IsLetter(runes[m.getCol()+1]) {
					continue
				}

				m.closeNode(newNode)
				parent.Children = append(parent.Children, *newNode)
				m.incrementColBy(len(typ)) // skip the closing symbols
				return nil
			}
			if m.isStartOfAnotherType() {
				return m.error("Nesting of types is not allowed!")
			}
		}
	}

	return errorFor(newNode, fmt.Sprintf("No closing %s found", typ))
}

func (m *MarkdownParser) parseTilda(parent *Node) *ParserError {
	return m.parseDefault("`", parent, true)
}

func (m *MarkdownParser) parseStar(parent *Node) *ParserError {
	return m.parseDefault("**", parent, true)
}

func (m *MarkdownParser) parseUnderscore(parent *Node) *ParserError {
	return m.parseDefault("_", parent, false)
}

func (m *MarkdownParser) parseText(parent *Node) {
	runes := m.curLineRunes()
	newNode := m.newNode(text)
	for ; m.getCol() < len(runes); m.incrementCol() {
		if m.getCol() != newNode.Pos.Col && m.isStartOfAnotherType() {
			m.closeNode(newNode)
			parent.Children = append(parent.Children, *newNode)
			return
		}
	}

	m.closeNode(newNode)
	parent.Children = append(parent.Children, *newNode)
	m.setCol(len(runes))
}

func (m *MarkdownParser) parsePreformatted(root *Node) *ParserError {
	lineIdx := m.getLine()
	line := m.input[lineIdx]
	if len(line) < 3 {
		return m.error("Invalid preformatted block")
	}

	m.incrementLine() // skip the opening ```
	newNode := m.newNode(preformatted)

	for ; m.getLine() < len(m.input); m.incrementLine() {
		line = m.input[m.getLine()]
		if strings.HasPrefix(line, "```") {
			if len(line) > 3 {
				return m.error("Invalid ending of the preformatted block")
			}
			newNode.Pos.EndLine = m.getLine() - 1
			newNode.Pos.EndCol = len(m.input[m.getLine()-1])
			newNode.IsClosed = true
			newNode.Val = m.getValOf(newNode.Pos)

			root.Children = append(root.Children, *newNode)
			m.setCol(3) // skip the closing ```
			return nil
		}
	}

	newNode.Pos.Line -= 1 // TODO: separate for block rules to avoid this hacks
	newNode.Pos.Col = 3
	return errorFor(newNode, "No closing ``` found")
}
