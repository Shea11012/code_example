package aes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAes(t *testing.T) {
	plain := "你好"

	cipherText := encrypt(plain)
	fmt.Printf("cipher: %s\n", cipherText)

	plainText := decrypt(cipherText)
	fmt.Printf("plainText: %s\n", plainText)

	require.Equal(t, plain, plainText)
}
