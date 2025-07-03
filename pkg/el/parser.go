package el

import (
	"errors"
)

// Expr : union of NameExpr, LambdaExpr
type Expr interface {
	String() string
	MustTypeExpr() // for type-safety every Expr must implement this
}

// NameExpr : a name, a number, a string, etc.
// e.g. name
type NameExpr string

func (e NameExpr) String() string {
	return string(e)
}

func (e NameExpr) MustTypeExpr() {}

// LambdaExpr : S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type LambdaExpr struct {
	Cmd  Name
	Args []Expr
}

func (e LambdaExpr) String() string {
	s := ""
	s += "("
	s += string(e.Cmd)
	for _, arg := range e.Args {
		s += " " + arg.String()
	}
	s += ")"
	return s
}

func (e LambdaExpr) MustTypeExpr() {}

func pop(tokenList []Token) ([]Token, Token, error) {
	if len(tokenList) == 0 {
		return nil, "", errors.New("empty token list")
	}
	return tokenList[1:], tokenList[0], nil
}

func Parse(tokenList []Token) (Expr, []Token, error) {
	tokenList, head, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}
	if head != "(" {
		return NameExpr(head), tokenList, nil
	}
	// parse LambdaExpr
	tokenList, cmd, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}
	if cmd == ")" {
		// empty LambdaExpr
		return nil, tokenList, nil
	}
	var arg Expr
	var argList []Expr
	for {
		arg, tokenList, err = Parse(tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		// detect end of LambdaExpr
		if nameExpr, ok := arg.(NameExpr); ok && string(nameExpr) == ")" {
			break
		}
		argList = append(argList, arg)
	}
	expr := LambdaExpr{
		Cmd:  Name(cmd),
		Args: argList,
	}
	return expr, tokenList, err
}
