package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

func GenerateECDHKeyPair() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func ComputeSharedSecret(privKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) []byte {
	x, _ := privKey.Curve.ScalarMult(pubKey.X, pubKey.Y, privKey.D.Bytes()) // S = dA * PB                             // Hash it for security
	sharedSecret := x.Bytes()
	hash := sha256.Sum256(sharedSecret)
	key := make([]byte, 32)
	copy(key, hash[:])

	return key
}
