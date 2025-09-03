package ast

// A very simple AST, each node is either an Expr or a list of Expr

// Expr - union of Leaf, Node
type Expr interface {
	String() string
}

// Leaf - a name, a number, a string, etc.
type Leaf string

func (e Leaf) String() string {
	return string(e)
}

// Node - S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type Node struct {
	Cmd  Expr
	Args []Expr
}

func (e Node) String() string {
	s := ""
	s += "("
	s += e.Cmd.String()
	for _, arg := range e.Args {
		s += " " + arg.String()
	}
	s += ")"
	return s
}
