package ast

type Token = string // TODO - type Token struct {Value string; Line int}

const (
	TokenBlockBegin Token = "("
	TokenBlockEnd   Token = ")"
	TokenSugarBegin Token = "{"
	TokenSugarEnd   Token = "}"
	TokenUnwrap     Token = "$"
	TokenTypeCast         = ":"
	TokenStringBeg  Token = "\""
	TokenStringEnd  Token = "\""
)

var SplitTokens = map[Token]struct{}{
	TokenBlockBegin: {},
	TokenBlockEnd:   {},
	TokenSugarBegin: {},
	TokenSugarEnd:   {},
	TokenUnwrap:     {},
	TokenTypeCast:   {},
}
