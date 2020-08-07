package hex

import "encoding/hex"

func HexEncodeString(ByteArray []byte) string {
	return hex.EncodeToString(ByteArray)
}

func HexDecodeString(DecodeString string) (string, error) {
	a, e := hex.DecodeString(DecodeString)
	if e != nil {
		return "", e
	}
	return string(a), nil
}
