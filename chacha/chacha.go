package chacha

import (
	"crypto/rand"

	"golang.org/x/crypto/chacha20"
)

func EncryptChaCha20(data []byte, key []byte) ([]byte, []byte) {
	nonce := make([]byte, chacha20.NonceSize)
	rand.Read(nonce)

	cipher, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	encrypted := make([]byte, len(data))
	cipher.XORKeyStream(encrypted, data)

	return encrypted, nonce
}

func DecryptChaCha20(encrypted []byte, key []byte, nonce []byte) ([]byte, error) {
	cipher, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	decrypted := make([]byte, len(encrypted))
	cipher.XORKeyStream(decrypted, encrypted)

	return decrypted, nil
}
