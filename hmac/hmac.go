package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

// ComputeHMAC generates an HMAC using the shared secret key
func ComputeHMAC(message string, key *big.Int) string {
	h := hmac.New(sha256.New, key.Bytes())
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
