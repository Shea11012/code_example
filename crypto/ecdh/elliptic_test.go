package ecdh

import (
	"crypto/elliptic"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEllipticECDH_GenerateSecret(t *testing.T) {
	e := NewEllipticECDH(elliptic.P256())
	aPriv, _ := e.GenerateEllipticKey()
	bPriv, _ := e.GenerateEllipticKey()

	aSecret := e.GenerateSecret(aPriv, bPriv.EllipticPublicKey)
	bSecret := e.GenerateSecret(bPriv, aPriv.EllipticPublicKey)
	require.Equal(t, aSecret, bSecret)
	fmt.Printf("aSecret: %s\nbSecret: %s\n", aSecret, bSecret)
}
