package parser

import (
	"fmt"
	"strings"
	"unicode"
)

type Token = string

func Tokenize(s string) []Token {
	return tokenizeWithSplitCharacters(s, map[rune]struct{}{
		'(': {},
		')': {},
		'*': {},
	})
}

func TokenizeWithInfixOperator(s string) []Token {
	// transpile first
	s = strings.ReplaceAll(s, "[[", " (list ")
	s = strings.ReplaceAll(s, "]]", " ) ")
	s = strings.ReplaceAll(s, "{", " (let ")
	s = strings.ReplaceAll(s, "}", " ) ")

	return tokenizeWithSplitCharacters(s, map[rune]struct{}{
		'(': {},
		')': {},
		'*': {},
		'[': {},
		']': {},
	})
}

func removeComments(str string) string {
	lines := strings.Split(str, "\n")
	var newLines []string
	for _, line := range lines {
		newLines = append(newLines, strings.Split(line, "//")[0])
	}
	return strings.Join(newLines, "\n")
}

func tokenizeWithSplitCharacters(str string, splitCharacters map[rune]struct{}) []Token {
	str = removeComments(str)

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
			} else if _, ok := splitCharacters[ch]; ok {
				// split characters are tokenized immediately
				flushBuffer()
				buffer += string(ch)
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
			panic(fmt.Sprintf("invalid state: %d", state))
		}
	}
	flushBuffer()
	return tokens
}
