package parser

import "el/ast"

func matchName(cond ast.Name) func(ast.Expr) bool {
	return func(arg ast.Expr) bool {
		if name, ok := arg.(ast.Name); ok {
			return string(cond) == string(name)
		}
		return false
	}
}
