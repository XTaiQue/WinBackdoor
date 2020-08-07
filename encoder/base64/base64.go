package base64

import "encoding/base64"

type Base64 struct {
	ByteArray string
}

func (b *Base64) EncodeToString() string {
	ByteArray := []byte(b.ByteArray)
	return base64.StdEncoding.EncodeToString(ByteArray)
}

func (b *Base64) DecodeToString() (string, error) {
	dec, e := base64.StdEncoding.DecodeString(b.ByteArray)
	if e != nil {
		return "", e
	}
	return string(dec), nil
}

func Base64EncodeToString(Bytearray []byte) string {
	return base64.StdEncoding.EncodeToString(Bytearray)
}

func Base64DecodeToString(String string) (string, error) {
	dec, e := base64.StdEncoding.DecodeString(String)
	if e != nil {
		return "", e
	}
	return string(dec), nil
}
