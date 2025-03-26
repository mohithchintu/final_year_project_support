package helpers

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"math/rand"
	"time"

	"github.com/mohithchintu/final_year_project_support/models"
)

// GenerateGroupKeyWithECDH generates the group key using ECDH and Shamir's Secret Sharing
func GenerateGroupKeyWithECDH(device *models.Device) error {
	// Use elliptic curve P-256 (or another curve suitable for IoT devices)
	curve := elliptic.P256()

	// Generate the ECC private key for the device
	priv, err := generateECCPrivateKey(curve)
	if err != nil {
		return fmt.Errorf("failed to generate ECC private key: %v", err)
	}

	// Store the ECC private key in the device
	device.PrivateKey = priv

	// Generate the public key for the device using the private key
	pubX, _ := curve.ScalarBaseMult(priv.Bytes()) // Only use the X-coordinate of the public key
	device.PublicKey = pubX

	// Ensure all peer devices also have their public keys set
	for _, peer := range device.Peers {
		// Generate the public key for each peer device if not already set
		if peer.PublicKey == nil {
			pubX, _ := curve.ScalarBaseMult(peer.PrivateKey.Bytes())
			peer.PublicKey = pubX
		}

		// Ensure the public key is valid before proceeding
		if !curve.IsOnCurve(peer.PublicKey, curve.Params().Gy) {
			return fmt.Errorf("peer device %s has an invalid public key", peer.ID)
		}
	}

	// Calculate the group key using ECDH and the other device's private key and shared information
	secret := big.NewInt(0)
	for _, peer := range device.Peers {
		// Derive the shared secret using this device's private key and the peer's public key
		sharedSecret, err := DeriveSharedSecret(curve, priv, peer.PublicKey)
		if err != nil {
			return fmt.Errorf("error deriving shared secret: %v", err)
		}

		// Combine the shared secrets
		secret.Add(secret, sharedSecret)
	}

	// Final group key derived from shared secrets (use modulo for consistency)
	secret.Mod(secret, big.NewInt(10000000999))
	device.GroupKey = secret

	return nil
}

// generateECCPrivateKey generates an ECC private key using the curve
func generateECCPrivateKey(curve elliptic.Curve) (*big.Int, error) {
	priv, _, _, err := elliptic.GenerateKey(curve, rand.New(rand.NewSource(time.Now().UnixNano())))
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECC private key: %v", err)
	}
	return new(big.Int).SetBytes(priv), nil
}

// DeriveSharedSecret derives a shared secret using ECDH (Elliptic Curve Diffie-Hellman)
func DeriveSharedSecret(curve elliptic.Curve, priv *big.Int, peerPubX *big.Int) (*big.Int, error) {
	// Ensure the peer's public key is valid before proceeding
	if peerPubX == nil {
		return nil, fmt.Errorf("peer's public key is nil")
	}

	// Compute the shared secret using the private key of this device and the public key of the peer device
	x, _ := curve.ScalarMult(peerPubX, curve.Params().Gy, priv.Bytes())

	// Return the x-coordinate as the shared secret
	return x, nil
}
