package ecdh

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

type EllipticECDH struct {
	curve elliptic.Curve
}

func NewEllipticECDH(curve elliptic.Curve) *EllipticECDH {
	return &EllipticECDH{
		curve: curve,
	}
}

type EllipticPublicKey struct {
	elliptic.Curve
	X, Y *big.Int
}

type EllipticPrivateKey struct {
	EllipticPublicKey
	D *big.Int
}

func (e *EllipticECDH) GenerateEllipticKey() (*EllipticPrivateKey, error) {
	priv, x, y, err := elliptic.GenerateKey(e.curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	publicKey := EllipticPublicKey{
		Curve: e.curve,
		X:     x,
		Y:     y,
	}

	return &EllipticPrivateKey{
		EllipticPublicKey: publicKey,
		D:                 big.NewInt(0).SetBytes(priv),
	}, nil
}

func (e *EllipticECDH) GenerateSecret(priv *EllipticPrivateKey, pub EllipticPublicKey) string {
	secret, _ := e.curve.ScalarMult(pub.X, pub.Y, priv.D.Bytes())
	return hex.EncodeToString(secret.Bytes())
}
