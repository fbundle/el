package el

import "strings"

func Transpile(s string) string {
	s = strings.ReplaceAll(s, "{", " (let ")
	s = strings.ReplaceAll(s, "}", " ) ")
	return s
}
