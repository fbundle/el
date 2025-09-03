package ast

import "strings"

// A very simple AST

// Expr - union of Name, Lambda
type Expr interface {
	String() string
}

// Name - a name, a number, a string, etc.
type Name string

func (e Name) String() string {
	return string(e)
}

// Lambda - S-expression - every enclosed by a pair of parentheses e.g. (cmd ...)
type Lambda []Expr

func (e Lambda) String() string {
	children := make([]string, 0, len(e))
	for _, child := range e {
		children = append(children, child.String())
	}
	return "(" + strings.Join(children, " ") + ")"
}
