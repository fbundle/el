package el

import "strings"

func Transpile(s string) string {
	s = strings.ReplaceAll(s, "]", " ) ")
	s = strings.ReplaceAll(s, "[", " (list ")
	s = strings.ReplaceAll(s, "}", " ) ")
	s = strings.ReplaceAll(s, "{", " (tail ")
	s = strings.ReplaceAll(s, ">", " ) ")
	s = strings.ReplaceAll(s, "<", " (sum ")
	return s
}
