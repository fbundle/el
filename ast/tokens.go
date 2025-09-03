package ast

type Token = string

const (
	TokenBlockBegin Token = "("
	TokenBlockEnd   Token = ")"
	TokenSugarBegin Token = "{"
	TokenSugarEnd   Token = "}"
	TokenUnwrap     Token = "*"
)

var SpecialTokens = map[Token]struct{}{
	TokenBlockBegin: {},
	TokenBlockEnd:   {},
	TokenSugarBegin: {},
	TokenSugarEnd:   {},
	TokenUnwrap:     {},
}
