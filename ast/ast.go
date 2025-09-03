package ast

// A very simple AST

// Node - union of Name, Function
type Node interface {
	String() string
}

// Name - a name, a number, a string, etc.
type Name string

func (e Name) String() string {
	return string(e)
}

// Function - S-expression - every enclosed by a pair of parentheses e.g. (cmd ...)
type Function struct {
	Cmd  Node
	Args []Node
}

func (e Function) String() string {
	s := ""
	s += "("
	s += e.Cmd.String()
	for _, arg := range e.Args {
		s += " " + arg.String()
	}
	s += ")"
	return s
}
