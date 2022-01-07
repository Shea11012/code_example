package ed25519

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// GetPrivatePEMBlock 使用ed25519生成私有证书
func GetPrivatePEMBlock(priKey ed25519.PrivateKey) (*pem.Block, error) {
	encodePrivateKey, err := x509.MarshalPKCS8PrivateKey(priKey)
	if err != nil {
		return nil, err
	}

	privatePem := &pem.Block{
		Type:  "ed15519 Private Key",
		Bytes: encodePrivateKey,
	}

	return privatePem, nil
}

func GetPublicPEMBlock(publicKey ed25519.PublicKey) (*pem.Block, error) {
	encodePublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	publicPem := &pem.Block{
		Type:  "Public Key",
		Bytes: encodePublicKey,
	}

	return publicPem, nil
}

func SavePemBlockToFile(pemBlock *pem.Block, filename string) error {
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = pem.Encode(outFile, pemBlock)
	if err != nil {
		return err
	}

	return nil
}
