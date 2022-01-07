package ecdh

import (
	"crypto/elliptic"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcdsaECDH_GenerateSecretKey(t *testing.T) {
	e := NewEcdsaECDH(elliptic.P256())
	aPriv, _ := e.GenerateKey()
	bPriv, _ := e.GenerateKey()

	aSecret := e.GenerateSecretKey(aPriv, bPriv.PublicKey)
	bSecret := e.GenerateSecretKey(bPriv, aPriv.PublicKey)
	require.Equal(t, aSecret, bSecret)
	fmt.Printf("aSecret: %s\nbSecret: %s\n", aSecret, bSecret)
}
