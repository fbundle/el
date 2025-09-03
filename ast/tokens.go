package ast

type Token = string

const (
	TokenBlockBegin Token = "("
	TokenBlockEnd   Token = ")"
	TokenSugarBegin Token = "{"
	TokenSugarEnd   Token = "}"
	TokenUnwrap     Token = "*"
	TokenStringBeg  Token = "\""
	TokenStringEnd  Token = "\""
)

var SplitTokens = map[Token]struct{}{
	TokenBlockBegin: {},
	TokenBlockEnd:   {},
	TokenSugarBegin: {},
	TokenSugarEnd:   {},
	TokenUnwrap:     {},
}
