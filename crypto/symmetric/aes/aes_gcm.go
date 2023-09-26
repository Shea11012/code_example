package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var secretKey = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"

func encrypt(plainText string) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)

	return string(cipherText)
}

func decrypt(cipherText string) string {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plainText, err := gcm.Open(nil, []byte(nonce), []byte(cipherText), nil)
	if err != nil {
		panic(err)
	}
	return string(plainText)
}
