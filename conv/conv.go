package conv

import (
	"fmt"
	"strconv"
)

func AsciiToInt(a byte) int {
	ascii := int(a)
	d := fmt.Sprintf("%d", ascii)
	i, _ := strconv.Atoi(d)
	return i
}

func IntToAscii(a int8) string {
	return fmt.Sprintf("%c", a)
}
