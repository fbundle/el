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

// LambdaExpr - S-expression - every enclosed by a pair of parentheses
// e.g. (cmd ...)
type LambdaExpr struct {
	Cmd  Expr
	Args []Expr
}

func (e LambdaExpr) String() string {
	s := ""
	s += "("
	s += e.Cmd.String()
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
			return LambdaExpr{
				Cmd:  argList[0],
				Args: argList[1:],
			}, tokenList, nil
		}
	} else {
		return NameExpr(head), tokenList, nil
	}
}

func ParseWithInfixOperator(tokenList []Token) (Expr, []Token, error) {
	tokenList, head, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}

	if head == "(" {
		// argList, tokenList, err := parseArgList("(", ")")
		argList, tokenList, err := parseUntilClose(tokenList, ")", ParseWithInfixOperator)
		if err != nil {
			return nil, tokenList, err
		}
		switch len(argList) {
		case 0:
			return nil, tokenList, errors.New("empty LambdaExpr")
		case 1:
			return argList[0], tokenList, nil
		default:
			return LambdaExpr{
				Cmd:  argList[0],
				Args: argList[1:],
			}, tokenList, nil
		}
	} else if head == "[" {
		argList, tokenList, err := parseUntilClose(tokenList, "]", ParseWithInfixOperator)
		if err != nil {
			return nil, tokenList, err
		}
		if len(argList)%2 == 0 {
			return nil, tokenList, errors.New("InfixOperator must have an odd number of arguments")
		}
		var parseInfixOperator func(argList []Expr) (Expr, error)
		parseInfixOperator = func(argList []Expr) (Expr, error) {
			// len argList is either 1 3 5 ...
			if len(argList) == 1 {
				return argList[0], nil
			}
			argList, cmdExpr, right := argList[:len(argList)-2], argList[len(argList)-2], argList[len(argList)-1]
			left, err := parseInfixOperator(argList)
			if err != nil {
				return nil, err
			}
			return LambdaExpr{
				Cmd:  cmdExpr,
				Args: []Expr{left, right},
			}, nil

		}

		expr, err := parseInfixOperator(argList)
		if err != nil {
			return nil, tokenList, err
		}
		return expr, tokenList, nil
	} else {
		return NameExpr(head), tokenList, nil
	}
}
