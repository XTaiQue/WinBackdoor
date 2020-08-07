package base85

import (
	"encoding/ascii85"
)

type Base85 struct {
	ByteArray string
}

func (b *Base85) EncodeToString() string {
	ByteArray := []byte(b.ByteArray)
	Arrayb85 := make([]byte, ascii85.MaxEncodedLen(len(b.ByteArray)))
	_ = ascii85.Encode(Arrayb85, ByteArray)
	return string(Arrayb85)
}

func (b *Base85) DecodeToString() string {
	ByteArray := []byte(b.ByteArray)
	Arrayb85 := make([]byte, ascii85.MaxEncodedLen(len(b.ByteArray)))
	nArrayDecodedBytes, _, _ := ascii85.Decode(Arrayb85, ByteArray, true)
	s := Arrayb85[:nArrayDecodedBytes]
	return string(s)
}

func Base85EncodeToString(Bytearray []byte) string {
	Arrayb85 := make([]byte, ascii85.MaxEncodedLen(len(Bytearray)))
	_ = ascii85.Encode(Arrayb85, Bytearray)
	return string(Arrayb85)
}

func Base85DecodeToString(Bytearray []byte) string {
	Arrayb85 := make([]byte, ascii85.MaxEncodedLen(len(Bytearray)))
	nArrayDecodedBytes, _, _ := ascii85.Decode(Arrayb85, Bytearray, true)
	return string(Arrayb85[:nArrayDecodedBytes])
}
