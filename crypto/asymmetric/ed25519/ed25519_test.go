package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Ed25519(t *testing.T) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	privatePem, err := GetPrivatePEMBlock(privateKey)
	require.NoError(t, err)

	publicPem, err := GetPublicPEMBlock(publicKey)
	require.NoError(t, err)

	err = SavePemBlockToFile(publicPem, "ed25519_pub.pem")
	require.NoError(t, err)

	err = SavePemBlockToFile(privatePem, "ed25519")
	require.NoError(t, err)
}
