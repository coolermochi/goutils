// crypt.
package goutils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// crypt
type crypt struct{}

// namespace
var Crypt = crypt{}

// type
type cryptType int

const (
	CryptPlain cryptType = iota
	CryptB64
	CryptB64URLSafe
)

// Encrypt.
func (crypt) Encrypt(typo cryptType, secret string, text string) string {

	secret = fmt.Sprintf("%016s", secret) // 0で16桁埋め

	str := encrypt(secret, text)

	switch typo {
	case CryptPlain:
		return str
	case CryptB64:
		return Str.EncodeB64(str)
	case CryptB64URLSafe:
		return Str.EncodeB64URLSafe(str)
	default:
		return str
	}
}

// Decrypt.
func (crypt) Decrypt(typo cryptType, secret string, data string) string {

	secret = fmt.Sprintf("%016s", secret) // 0で16桁埋め

	switch typo {
	case CryptPlain:
	case CryptB64:
		data = Str.DecodeB64(data)
	case CryptB64URLSafe:
		data = Str.DecodeB64URLSafe(data)
	}

	return decrypt(secret, data)
}

// encrypt.
func encrypt(secret string, text string) string {
	t := []byte(text)
	key := []byte(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error NewCipher secret:%s text:%s %s\n", secret, text, err.Error())
		return ""
	}

	cipherText := make([]byte, aes.BlockSize+len(t))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("error ReadFull secret:%s text:%s %s\n", secret, text, err.Error())
	}

	// Encrypt
	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], t)

	return string(cipherText)
}

// decrypt.
func decrypt(secret string, data string) string {
	cipherText := []byte(data)
	key := []byte(secret)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("error NewCipher secret:%s data:%s %s\n", secret, data, err.Error())
		return ""
	}
	// Decrypt
	decryptedText := make([]byte, len(cipherText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, cipherText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, cipherText[aes.BlockSize:])

	return string(decryptedText)
}
