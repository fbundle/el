package parser

import (
	"el/ast"
	"errors"

	"github.com/fbundle/lab_public/lab/go_util/pkg/adt"
)

func pop(tokenList []Token) ([]Token, Token, error) {
	if len(tokenList) == 0 {
		return nil, "", errors.New("empty token list")
	}
	return tokenList[1:], tokenList[0], nil
}

type Parser = func(tokenList []Token) (ast.Expr, []Token, error)

func parseUntil(parser Parser, stopCond func(ast.Expr) bool, tokenList []Token) ([]ast.Expr, []Token, error) {
	var arg ast.Expr
	var err error
	var argList []ast.Expr
	for {
		arg, tokenList, err = parser(tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		if stopCond(arg) {
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

	switch head {
	case TokenBlockBegin:
		// parse until seeing `)`
		argList, tokenList, err := parseUntil(Parse, matchName(ast.Name(TokenBlockEnd)), tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		return ast.Lambda(argList), tokenList, nil
	case TokenSugarBegin:
		// parse until seeing `}`
		argList, tokenList, err := parseUntil(Parse, matchName(ast.Name(TokenSugarEnd)), tokenList)
		if err != nil {
			return nil, tokenList, err
		}
		expr, err := processSugar(argList)
		return expr, tokenList, err
	default:
		return ast.Name(head), tokenList, nil
	}
}

// processSugar - handles both arithmetic infix and lambda syntax
// {1 + 2 + 3} -> (add (add 1 2) 3)
// {x y => (add x y)} -> (lambda x y (add x y))
func processSugar(argList []ast.Expr) (ast.Expr, error) {
	if len(argList) == 0 {
		return ast.Lambda(nil), nil
	}
	if len(argList) == 1 {
		return argList[0], nil
	}
	var arrow ast.Name
	if ok := adt.Cast[ast.Name](argList[len(argList)-2]).Unwrap(&arrow); ok && string(arrow) == "=>" {
		// arrow function syntax: {x y => expr}
		paramList := argList[:len(argList)-2]
		body := argList[len(argList)-1]
		lambdaArgList := []ast.Expr{
			ast.Name("lambda"),
		}
		lambdaArgList = append(lambdaArgList, paramList...)
		lambdaArgList = append(lambdaArgList, body)

		return ast.Lambda(lambdaArgList), nil
	}

	// No arrow function found, process as regular infix
	argList, cmd, right := argList[:len(argList)-2], argList[len(argList)-2], argList[len(argList)-1]
	left, err := processSugar(argList)
	if err != nil {
		return nil, err
	}
	return ast.Lambda([]ast.Expr{cmd, left, right}), nil
}
