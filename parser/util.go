package parser

import "el/ast"

const (
	TokenBlockBegin Token = "("
	TokenBlockEnd   Token = ")"
	TokenSugarBegin Token = "{"
	TokenSugarEnd   Token = "}"
	TokenUnwrap     Token = "*"
)

var specialTokens = map[Token]struct{}{
	TokenBlockBegin: {},
	TokenBlockEnd:   {},
	TokenSugarBegin: {},
	TokenSugarEnd:   {},
	TokenUnwrap:     {},
}

var specialChars = map[rune]struct{}{
	'(': {},
	')': {},
	'{': {},
	'}': {},
	'*': {},
}

func matchName(cond ast.Name) func(ast.Expr) bool {
	return func(arg ast.Expr) bool {
		if name, ok := arg.(ast.Name); ok {
			return string(cond) == string(name)
		}
		return false
	}
}
