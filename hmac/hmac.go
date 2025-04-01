package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
)

func GenerateHMAC(data string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return h.Sum(nil)
}

func VerifyHMAC(data string, key []byte, hmacValue []byte) bool {
	expectedHMAC := GenerateHMAC(data, key)
	return hmac.Equal(expectedHMAC, hmacValue)
}
