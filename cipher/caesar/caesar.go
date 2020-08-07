package caesar

import (
	"strings"
)

func cipher(r rune, shift int) rune {
	s := int(r) + shift

	if s > 'z' {
		return rune(s - 26)
	} else if s < 'a' {
		return rune(s + 26)
	}

	return rune(s)
}

func Caesar(text string, shift int) string {
	caes := strings.Map(func(r rune) rune {
		return cipher(r, shift)
	}, text)

	return caes
}
