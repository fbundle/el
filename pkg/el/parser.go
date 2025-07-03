package el

import (
	"errors"
	"slices"
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

type Parser = func(tokenList []Token) (Expr, []Token, error)

func parseUntilClose(tokenList []Token, close Token, parseOnce Parser) ([]Expr, []Token, error) {
	var arg Expr
	var err error
	var argList []Expr
	for {
		arg, tokenList, err = parseOnce(tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		// detect end of LambdaExpr
		if nameExpr, ok := arg.(NameExpr); ok && string(nameExpr) == close {
			break
		}
		argList = append(argList, arg)
	}
	return argList, tokenList, nil
}

func Parse(tokenList []Token) (Expr, []Token, error) {
	tokenList, head, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}

	if head == "(" {
		// argList, tokenList, err := parseArgList("(", ")")
		argList, tokenList, err := parseUntilClose(tokenList, ")", Parse)
		if err != nil {
			return nil, tokenList, err
		}
		switch len(argList) {
		case 0:
			return nil, tokenList, errors.New("empty LambdaExpr")
		case 1:
			return argList[0], tokenList, nil
		default:
			cmd, ok := argList[0].(NameExpr)
			if !ok {
				return nil, tokenList, errors.New("LambdaExpr must start with a name")
			}
			return LambdaExpr{
				Cmd:  Name(cmd),
				Args: argList[1:],
			}, tokenList, nil
		}
	} else {
		return NameExpr(head), tokenList, nil
	}
}

func ParseWithInplaceOperator(tokenList []Token) (Expr, []Token, error) {
	tokenList, head, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}

	if head == "(" {
		// argList, tokenList, err := parseArgList("(", ")")
		argList, tokenList, err := parseUntilClose(tokenList, ")", ParseWithInplaceOperator)
		if err != nil {
			return nil, tokenList, err
		}
		switch len(argList) {
		case 0:
			return nil, tokenList, errors.New("empty LambdaExpr")
		case 1:
			return argList[0], tokenList, nil
		default:
			cmd, ok := argList[0].(NameExpr)
			if !ok {
				return nil, tokenList, errors.New("LambdaExpr must start with a name")
			}
			return LambdaExpr{
				Cmd:  Name(cmd),
				Args: argList[1:],
			}, tokenList, nil
		}
	} else if head == "[" {
		argList, tokenList, err := parseUntilClose(tokenList, "]", ParseWithInplaceOperator)
		if err != nil {
			return nil, tokenList, err
		}
		if len(argList)%2 == 0 {
			return nil, tokenList, errors.New("InplaceOperator must have an odd number of arguments")
		}
		var parseInplaceOperator func(argList []Expr) (Expr, error)
		parseInplaceOperator = func(argList []Expr) (Expr, error) {
			if len(argList) == 0 {
				return nil, errors.New("empty InplaceOperator")
			}
			if len(argList) == 1 {
				return argList[0], nil
			}
			if _, ok := argList[1].(NameExpr); !ok {
				return nil, errors.New("InplaceOperator must have a operator")
			}
			right, err := parseInplaceOperator(argList[2:])
			if err != nil {
				return nil, err
			}
			return LambdaExpr{
				Cmd:  Name(argList[1].(NameExpr)),
				Args: []Expr{argList[0], right},
			}, nil
		}
		slices.Reverse(argList)
		expr, err := parseInplaceOperator(argList)
		if err != nil {
			return nil, tokenList, err
		}
		return expr, tokenList, nil
	} else {
		return NameExpr(head), tokenList, nil
	}
}
