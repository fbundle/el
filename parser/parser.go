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

func parseArgList(parser Parser, stopExpr ast.Name, tokenList []Token) ([]ast.Expr, []Token, error) {
	isStopExpr := func(arg ast.Expr) bool {
		if name, ok := arg.(ast.Name); ok {
			return string(name) == string(stopExpr)
		}
		return false
	}
	var arg ast.Expr
	var err error
	var argList []ast.Expr
	for {
		arg, tokenList, err = parser(tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		if isStopExpr(arg) {
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

	if head != "(" {
		return ast.Name(head), tokenList, nil
	}
	// parse until seeing `)`
	argList, tokenList, err := parseArgList(Parse, ")", tokenList)
	if err != nil {
		return nil, tokenList, err
	}
	return ast.Lambda(argList), tokenList, nil
}

func ParseWithInfixOperator(tokenList []Token) (ast.Expr, []Token, error) {
	tokenList, head, err := pop(tokenList)
	if err != nil {
		return nil, tokenList, err
	}

	switch head {
	case "(":
		// parse until seeing `)`
		argList, tokenList, err := parseArgList(ParseWithInfixOperator, ")", tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		return ast.Lambda(argList), tokenList, nil
	case "{":
		// parse until seeing `}`
		argList, tokenList, err := parseArgList(ParseWithInfixOperator, "}", tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		expr, err := processInfix(argList)
		return expr, tokenList, err

	default:
		return ast.Name(head), tokenList, nil
	}
}

// processInfix - [1 + 2 + 3] -> (add (add 1 2) 3)
func processInfix(argList []ast.Expr) (ast.Expr, error) {
	if len(argList) == 0 {
		return ast.Lambda(nil), nil
	}
	if len(argList) == 1 {
		return argList[0], nil
	}

	argList, cmd, right := argList[:len(argList)-2], argList[len(argList)-2], argList[len(argList)-1]
	left, err := processInfix(argList)
	if err != nil {
		return nil, err
	}
	return ast.Lambda([]ast.Expr{cmd, left, right}), nil

}
