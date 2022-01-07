package ecdh

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

type EcdsaECDH struct {
	curve elliptic.Curve
}

func NewEcdsaECDH(curve elliptic.Curve) *EcdsaECDH {
	return &EcdsaECDH{
		curve: curve,
	}
}

func (e *EcdsaECDH) GenerateKey() (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(e.curve, rand.Reader)
	if err != nil {
		return nil, nil
	}

	return priv, err
}

func (e *EcdsaECDH) GenerateSecretKey(priv *ecdsa.PrivateKey, pub ecdsa.PublicKey) string {
	secret, _ := pub.Curve.ScalarMult(pub.X, pub.Y, priv.D.Bytes())
	return hex.EncodeToString(secret.Bytes())
}
