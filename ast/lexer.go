package ast

import (
	"fmt"
	"strings"
	"unicode"
)

type Token = string

func Tokenize(s string) []Token {
	return tokenize(s,
		removeComment("#"),
		splitString([]string{
			"(",
			")",
			"*",
		}),
	)
}

func TokenizeWithInfixOperator(s string) []Token {
	return tokenize(s,
		removeComment("#"),
		mapping(map[string]string{
			"[[": " (list ",
			"]]": " ) ",
			"{":  " (let ",
			"}":  " ) ",
		}),
		splitString([]string{
			"(",
			")",
			"*",
			"[",
			"]",
		}),
	)
}

type preprocessor func(string) string

var mapping = func(stringMap map[string]string) preprocessor {
	return func(str string) string {
		for k, v := range stringMap {
			str = strings.ReplaceAll(str, k, v)
		}
		return str
	}
}

var removeComment = func(sep string) preprocessor {
	return func(str string) string {
		lines := strings.Split(str, "\n")
		var newLines []string
		for _, line := range lines {
			newLines = append(newLines, strings.SplitN(line, sep, 2)[0])
		}
		return strings.Join(newLines, "\n")
	}
}

var splitString = func(sepString []string) preprocessor {
	return func(str string) string {
		normalize := func(s string) string {
			return strings.Join(strings.Fields(s), " ")
		}
		str = normalize(str)
		for {
			str1 := str
			for _, s := range sepString {
				str1 = strings.ReplaceAll(str1, s, " "+s+" ")
			}
			str1 = normalize(str1)
			if str1 == str {
				break
			}
			str = str1
		}
		return str
	}
}

func tokenize(str string, pList ...preprocessor) []Token {
	for _, p := range pList {
		str = p(str)
	}

	const (
		STATE_OUTSTRING = iota
		STATE_INSTRING
		STATE_INSTRING_ESCAPE
	)

	var tokens []Token
	state := STATE_OUTSTRING
	buffer := ""
	flushBuffer := func() {
		if len(buffer) > 0 {
			tokens = append(tokens, buffer)
		}
		buffer = ""
	}
	for _, ch := range str {
		switch state {
		case STATE_OUTSTRING: // outside string
			if unicode.IsSpace(ch) {
				// flush buffer if seeing whitespace
				flushBuffer()
			} else if ch == '"' {
				// enter string mode
				flushBuffer()
				buffer += string(ch)
				state = STATE_INSTRING
			} else {
				buffer += string(ch)
			}
		case STATE_INSTRING:
			if ch == '\\' {
				buffer += string(ch)
				state = STATE_INSTRING_ESCAPE
			} else if ch == '"' {
				// exit string mode
				buffer += string(ch)
				flushBuffer()
				state = STATE_OUTSTRING
			} else {
				buffer += string(ch)
			}
		case STATE_INSTRING_ESCAPE:
			buffer += string(ch)
			state = STATE_INSTRING
		default:
			panic(fmt.Sprintf("unreachable state: %d", state))
		}
	}
	flushBuffer()
	return tokens
}
