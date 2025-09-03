package ast

// A very simple AST, each node is either an Expr or a list of Expr

// Expr : union of Atom, SExpr
type Expr interface {
	String() string
}

// Atom - a name, a number, a string, etc.
type Atom string

func (e Atom) String() string {
	return string(e)
}

// SExpr - S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type SExpr struct {
	Cmd  Expr
	Args []Expr
}

func (e SExpr) String() string {
	s := ""
	s += "("
	s += e.Cmd.String()
	for _, arg := range e.Args {
		s += " " + arg.String()
	}
	s += ")"
	return s
}
