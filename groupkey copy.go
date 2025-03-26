// package helpers
// import (
// 	"crypto/elliptic"
// 	//"crypto/rand"
// 	"fmt"
// 	"math/rand"
// 	"math/big"
// 	//"math/rand"
// 	"time"

// 	"github.com/mohithchintu/final_year_project_support/models"
// )

// // GenerateGroupKeyWithECDH generates the group key using ECDH and Shamir's Secret Sharing
// func GenerateGroupKeyWithECDH(device *models.Device) error {
// 	// Use elliptic curve P-256 (or another curve suitable for IoT devices)
// 	curve := elliptic.P256()

// 	// Generate the ECC private key for the device
// 	priv, err := generateECCPrivateKey(curve)
// 	if err != nil {
// 		return fmt.Errorf("failed to generate ECC private key: %v", err)
// 	}

// 	// Store the ECC private key in the device
// 	device.PrivateKey = priv

// 	// Calculate the group key using ECDH and the other device's private key and shared information
// 	secret := big.NewInt(0)
// 	for _, peer := range device.Peers {
// 		peerPrivateKey := peer.PrivateKey

// 		// Derive the shared secret (ECDH) with the peer's private key and this device's private key
// 		sharedSecret, err := DeriveSharedSecret(curve, priv, peerPrivateKey)
// 		if err != nil {
// 			return fmt.Errorf("error deriving shared secret: %v", err)
// 		}

// 		// Combine the shared secrets
// 		secret.Add(secret, sharedSecret)
// 	}

// 	// Final group key derived from shared secrets (use modulo for consistency)
// 	secret.Mod(secret, big.NewInt(10000000999))
// 	device.GroupKey = secret

// 	return nil
// }

// // generateECCPrivateKey generates an ECC private key using the curve
// func generateECCPrivateKey(curve elliptic.Curve) (*big.Int, error) {
// 	priv, _, _, err := elliptic.GenerateKey(curve, rand.New(rand.NewSource(time.Now().UnixNano())))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate ECC private key: %v", err)
// 	}
// 	return new(big.Int).SetBytes(priv), nil
// }

// // deriveSharedSecret derives a shared secret using ECDH (Elliptic Curve Diffie-Hellman)
// func DeriveSharedSecret(curve elliptic.Curve, priv *big.Int, peerPriv *big.Int) (*big.Int, error) {
// 	// Calculate the shared secret by performing ECDH on the curve using the private key of both devices
// 	x, _ := curve.ScalarMult(curve.Params().Gx, curve.Params().Gy, priv.Bytes())

// 	// Return the x-coordinate as the shared secret
// 	return x, nil
// }
