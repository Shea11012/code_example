package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type GCM struct {
	key []byte
	gcm cipher.AEAD
}

func NewGCM(key []byte) (*GCM, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	return &GCM{
		key: key,
		gcm: gcm,
	}, nil
}

func (g *GCM) Encrypt(message []byte) ([]byte, error) {
	nonce := make([]byte, g.gcm.NonceSize())
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	cipherText := g.gcm.Seal(nil, nonce, message, nil)
	return append(nonce, cipherText...), nil
}

func (g *GCM) Decrypt(cipherText []byte) ([]byte, error) {
	nonce := cipherText[:g.gcm.NonceSize()]
	cipherText = cipherText[g.gcm.NonceSize():]
	return g.gcm.Open(nil, nonce, cipherText, nil)
}
