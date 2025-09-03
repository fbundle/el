package ast

import "strings"

// A very simple AST

type Token = string

// Expr - union of Leaf, Node
type Expr interface {
	String() string
}

// Leaf - a name, a number, a string, etc.
type Leaf Token

func (e Leaf) String() string {
	return string(e)
}

// Node - S-expression - every enclosed by a pair of parentheses e.g. (cmd ...)
type Node []Expr

func (e Node) String() string {
	children := make([]string, 0, len(e))
	for _, child := range e {
		children = append(children, child.String())
	}
	return "(" + strings.Join(children, " ") + ")"
}
