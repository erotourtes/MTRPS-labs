package common

type Parser interface {
	Parse() error
	GetNodes() []Node
}

type Node struct {
	Val      string
	Type     string
	Children []Node
	Pos      *Pos
	IsClosed bool
}

type Pos struct {
	Line    int
	Col     int
	EndLine int
	EndCol  int
}

const (
	Bold         = "b"
	Text         = "text"
	LineBreak    = "lineBreak"
	Italic       = "i"
	Monospace    = "tt"
	Preformatted = "pre"
)
