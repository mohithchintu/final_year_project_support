package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

func ComputeHMAC(message string, key *big.Int) string {
	h := hmac.New(sha256.New, key.Bytes())
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
