package aes

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGCM(t *testing.T) {
	key := "/NLMd0awkwimVLec8XJAa3mS18gxk9tpI8vc//1CGEI="
	data, err := base64.StdEncoding.DecodeString(key)
	require.NoError(t, err)
	aesGCM, err := NewGCM(data)
	require.NoError(t, err)

	text := "hello"

	encryptText, err := aesGCM.Encrypt([]byte(text))
	require.NoError(t, err)

	plainText, err := aesGCM.Decrypt(encryptText)
	require.NoError(t, err)

	require.Equal(t, text, string(plainText))
}

func BenchmarkGCM_Encrypt(b *testing.B) {
	key := "/NLMd0awkwimVLec8XJAa3mS18gxk9tpI8vc//1CGEI="
	data, _ := base64.StdEncoding.DecodeString(key)
	aesGCM, _ := NewGCM(data)
	text := "hello"
	for i := 0; i < b.N; i++ {
		_, _ = aesGCM.Encrypt([]byte(text))
	}
}
