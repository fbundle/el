package parser

import (
	"el/ast"
	"errors"
)

func pop(tokenList []Token) ([]Token, Token, error) {
	if len(tokenList) == 0 {
		return nil, "", errors.New("empty token list")
	}
	return tokenList[1:], tokenList[0], nil
}

type Parser = func(tokenList []Token) (ast.Expr, []Token, error)

func parseUntilClose(tokenList []Token, close Token, parseOnce Parser) ([]ast.Expr, []Token, error) {
	var arg ast.Expr
	var err error
	var argList []ast.Expr
	for {
		arg, tokenList, err = parseOnce(tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		// detect end of Node
		if nameExpr, ok := arg.(ast.Leaf); ok && string(nameExpr) == close {
			break
		}
		argList = append(argList, arg)
	}
	return argList, tokenList, nil
}

func Parse(tokenList []Token) (ast.Expr, []Token, error) {
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
			return nil, tokenList, errors.New("empty Node")
		case 1:
			return argList[0], tokenList, nil
		default:
			return ast.Node(argList), tokenList, nil
		}
	} else {
		return ast.Leaf(head), tokenList, nil
	}
}

func ParseWithInfixOperator(tokenList []Token) (ast.Expr, []Token, error) {
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
		if len(argList) == 0 {
			return nil, tokenList, errors.New("empty Node")
		}
		return ast.Node(argList), tokenList, nil
	} else if head == "[" {
		argList, tokenList, err := parseUntilClose(tokenList, "]", ParseWithInfixOperator)
		if err != nil {
			return nil, tokenList, err
		}
		if len(argList)%2 == 0 {
			return nil, tokenList, errors.New("InfixOperator must have an odd number of arguments")
		}
		var parseInfixOperator func(argList []ast.Expr) (ast.Expr, error)
		parseInfixOperator = func(argList []ast.Expr) (ast.Expr, error) {
			// len argList is either 1 3 5 ...
			if len(argList) == 1 {
				return argList[0], nil
			}
			argList, cmdExpr, right := argList[:len(argList)-2], argList[len(argList)-2], argList[len(argList)-1]
			left, err := parseInfixOperator(argList)
			if err != nil {
				return nil, err
			}

			return ast.Node([]ast.Expr{cmdExpr, left, right}), nil

		}

		expr, err := parseInfixOperator(argList)
		if err != nil {
			return nil, tokenList, err
		}
		return expr, tokenList, nil
	} else {
		return ast.Leaf(head), tokenList, nil
	}
}
