package renderer

type Parser interface {
	Parse() error
	GetNodes() []Node
}

type Node struct {
	Val  string
	Type string
}
