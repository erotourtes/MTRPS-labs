package common

type Parser interface {
	Parse() error
	GetNodes() []Node
}

type Node struct {
	Val  string
	Type string
}

const (
	Bold         = "b"
	Text         = "text"
	LineBreak    = "lineBreak"
	Italic       = "i"
	Monospace    = "tt"
	Preformatted = "pre"
)
