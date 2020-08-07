package xor

func Xor(String string, key string) (xored string) {
	for x := 0; x < len(String); x++ {
		xored += string(String[x] ^ key[x%len(key)])
	}
	return xored
}

func XorLogical(a int, b int) int {
	if a != b {
		return 1
	} else {
		return 0
	}
}
