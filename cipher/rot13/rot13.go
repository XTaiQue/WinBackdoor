package rot13

import "strings"

func cipher(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'a' && r <= 'z' {
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}

func Rot13(text string) string {
	enc := strings.Map(cipher, text)
	return enc
}
