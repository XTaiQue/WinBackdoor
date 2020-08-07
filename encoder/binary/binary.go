package binary

import (
	"bytes"
	"fmt"
)

func BinaryBytes(s string) string {
	var buffer bytes.Buffer
	for x := 0; x < len(s); x++ {
		fmt.Fprintf(&buffer, "%b", s[x]) // b for binary format
	}

	return fmt.Sprintf("%s", buffer.Bytes())
}
