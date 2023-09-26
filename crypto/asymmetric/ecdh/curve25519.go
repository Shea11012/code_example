// Package ecdh 椭圆曲线数字密钥交换算法，基于ECC的curve25519
package ecdh

import (
	"crypto/ecdh"
	"crypto/rand"
	"fmt"
)

func exchange() {
	// 可以更换为其他的ecc算法
	x25519 := ecdh.X25519()
	aliceKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	alicePubKey, err := x25519.NewPublicKey(aliceKey.PublicKey().Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Printf("=====alice\naliceKey: %x\nalicePubKey: %x\n", aliceKey.Bytes(), alicePubKey.Bytes())

	bobKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	bobPubKey, err := x25519.NewPublicKey(bobKey.PublicKey().Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Printf("====bob\nbobKey: %x\nbobPubKey: %x\n", bobKey.Bytes(), bobPubKey.Bytes())

	aliceShareKey, err := aliceKey.ECDH(bobKey.PublicKey())
	if err != nil {
		panic(err)
	}

	bobShareKey, err := bobKey.ECDH(aliceKey.PublicKey())
	if err != nil {
		panic(err)
	}

	fmt.Printf("=====shareKey\naliceShareKey: %x\nbobShareKey: %x\n", aliceShareKey, bobShareKey)
}
