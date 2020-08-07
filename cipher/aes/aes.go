package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	wbde "github.com/Unam3dd/WinBackdoor/winbderror"
)

type AES struct {
	Key       string
	HexFormat bool
}

func (a *AES) CBCEncrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CBC mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}

	if len(plaintext) != len(a.Key) {
		return "", wbde.New("error : [!] the plain text must be equal of size of key")
	}

	c := make([]byte, aes.BlockSize+len(plaintext))
	initialization_vector := c[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initialization_vector)

	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, initialization_vector)
	mode.CryptBlocks(c[aes.BlockSize:], plaintext)

	if a.HexFormat != false {
		return fmt.Sprintf("%X", c), nil
	}

	return fmt.Sprintf("%s", c), nil
}

func (a *AES) CBCDecrypt(ciphertext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CBC mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}

	if len(ciphertext) < aes.BlockSize {
		return "", wbde.New("error : ciphertext is too short in relation to BlockSize")
	}

	initialization_vector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", wbde.New("error : ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, initialization_vector)
	mode.CryptBlocks(ciphertext, ciphertext)

	if a.HexFormat != false {
		return fmt.Sprintf("%X", ciphertext), nil
	}

	return fmt.Sprintf("%s", ciphertext), nil
}

func (a *AES) CTREncrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CTR mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	initialization_vector := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initialization_vector)

	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, initialization_vector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	if a.HexFormat != false {
		return fmt.Sprintf("%X", ciphertext), nil
	}
	return fmt.Sprintf("%s", ciphertext), nil
}

func (a *AES) CTRDecrypt(ciphertext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CTR mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, ciphertext[:aes.BlockSize])
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])

	if a.HexFormat != false {
		return fmt.Sprintf("%X", plaintext), nil
	}

	return fmt.Sprintf("%s", plaintext), nil
}

func (a *AES) CFBEncrypt(plaintext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CFB mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	initialization_vector := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, initialization_vector)

	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, initialization_vector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	if a.HexFormat != false {
		return fmt.Sprintf("%X", ciphertext), nil
	}

	return fmt.Sprintf("%s", ciphertext), nil
}

func (a *AES) CFBDecrypt(ciphertext []byte) (string, error) {
	block, err := aes.NewCipher([]byte(a.Key))

	if err != nil {
		es := fmt.Sprintf("%s", err)
		return "", wbde.New(es + " => in CFB mode the key must be a key of size 16,32,64 bytes, because 16x8=AES-128, 32x8=AES-196, 64x8=AES-256")
	}

	if len(ciphertext) < aes.BlockSize {
		return "", wbde.New("error : ciphertext is too short in relation to BlockSize")
	}

	initialization_vector := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, initialization_vector)
	stream.XORKeyStream(ciphertext, ciphertext)

	if a.HexFormat != false {
		return fmt.Sprintf("%X", ciphertext), nil
	}

	return fmt.Sprintf("%s", ciphertext), nil
}
