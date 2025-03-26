package models

import "math/big"

// Share represents a share of the secret in Shamir's Secret Sharing scheme
type Share struct {
	X *big.Int
	Y *big.Int
}
