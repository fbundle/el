package expr

// A very simple AST, each node is either an Expr or a list of Expr

// Expr : union of Name, Lambda
type Expr interface {
	MustString() string
	MustTypeExpr() // for type-safety every Expr must implement this
}

// Name - a name, a number, a string, etc.
type Name string

func (e Name) MustString() string {
	return string(e)
}

func (e Name) MustTypeExpr() {}

// Lambda - S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type Lambda struct {
	Cmd  Expr
	Args []Expr
}

func (e Lambda) MustString() string {
	s := ""
	s += "("
	s += e.Cmd.MustString()
	for _, arg := range e.Args {
		s += " " + arg.MustString()
	}
	s += ")"
	return s
}

func (e Lambda) MustTypeExpr() {}
