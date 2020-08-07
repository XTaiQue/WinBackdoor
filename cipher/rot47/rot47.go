package rot47

import "strings"

func cipher(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		if r >= 'M' {
			return r - 47
		} else {
			return r + 47
		}
	} else if r >= 'a' && r <= 'z' {
		if r >= 'm' {
			return r - 47
		} else {
			return r + 47
		}
	}
	return r
}

func Rot47(text string) string {
	enc := strings.Map(cipher, text)
	return enc
}
