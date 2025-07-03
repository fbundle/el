package ast

// A very simple AST, each node is either an Expr or a list of Expr

// Expr : union of Name, Lambda
type Expr interface {
	String() string
	MustTypeExpr() // for type-safety every Expr must implement this
}

// Name - a name, a number, a string, etc.
type Name string

func (e Name) String() string {
	return string(e)
}

func (e Name) MustTypeExpr() {}

// Lambda - S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type Lambda struct {
	Cmd  Expr
	Args []Expr
}

func (e Lambda) String() string {
	s := ""
	s += "("
	s += e.Cmd.String()
	for _, arg := range e.Args {
		s += " " + arg.String()
	}
	s += ")"
	return s
}

func (e Lambda) MustTypeExpr() {}
