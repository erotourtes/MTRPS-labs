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
	BoldT         = "bold"
	TextT         = "text"
	LineBreakT    = "lineBreak"
	ItalicT       = "italic"
	MonospaceT    = "monospace"
	PreformattedT = "preformatted"
)
