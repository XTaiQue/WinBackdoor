package base32

import "encoding/base32"

type Base32 struct {
	ByteArray string
}

func (b *Base32) EncodeToString() string {
	ByteArray := []byte(b.ByteArray)
	return base32.StdEncoding.EncodeToString(ByteArray)
}

func (b *Base32) DecodeToString() (string, error) {
	dec, e := base32.StdEncoding.DecodeString(b.ByteArray)
	if e != nil {
		return "", e
	}
	return string(dec), nil
}

func Base32EncodeToString(Bytearray []byte) string {
	return base32.StdEncoding.EncodeToString(Bytearray)
}

func Base32DecodeToString(String string) (string, error) {
	dec, e := base32.StdEncoding.DecodeString(String)
	if e != nil {
		return "", e
	}
	return string(dec), nil
}
