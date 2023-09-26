// Package ecdsa 椭圆曲线数字签名算法。用于对数据创建数字签名，
// 以保障信息在传递和使用过程中的完整性、真实性和不可篡改。
package ecdsa

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func sign() {
	msg := "我是输入内容"

	// 此处可以更换为其他的ecc算法
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)

	// 可以更换为其他的hash函数
	digest := sha256.Sum256([]byte(msg))

	sig := ed25519.Sign(priv, digest[:])

	fmt.Printf("msg: %s\nhash: %x\n", msg, digest)
	fmt.Printf("\n===key===\n")
	fmt.Printf("pub: %x\n", pub)
	fmt.Printf("priv key: (priv=%x,pub=%x)\n", priv[:32], priv[32:])
	fmt.Printf("\n===signature===\n")
	fmt.Printf("signature: (R=%x,s=%x)\n", sig[:32], sig[32:])

	equal := ed25519.Verify(pub, digest[:], sig)
	fmt.Printf("digest equal sig: %t\n", equal)
}
