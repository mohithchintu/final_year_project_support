package sss

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/models"
)

// ReconstructPolynomialWithECDH reconstructs the polynomial secret using ECDH-derived shared secrets
func ReconstructPolynomialWithECDH(shares []*models.Share, curve elliptic.Curve, device *models.Device) (*big.Int, error) {
	if len(shares) < device.Threshold {
		return nil, fmt.Errorf("at least %d shares are required for reconstruction", device.Threshold)
	}

	secret := big.NewInt(0)
	modulus := big.NewInt(10000000999)

	for i, share := range shares {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)

		for j, otherShare := range shares {
			if i != j {
				peerPriv := otherShare.Y
				_, err := helpers.DeriveSharedSecret(curve, device.PrivateKey, peerPriv)
				if err != nil {
					return nil, fmt.Errorf("error in deriving shared secret: %v", err)
				}

				numerator.Mul(numerator, new(big.Int).Neg(share.X))
				denominator.Mul(denominator, new(big.Int).Sub(share.X, otherShare.X))
			}
		}

		denominatorInverse := new(big.Int).ModInverse(denominator, modulus)
		if denominatorInverse == nil {
			return nil, fmt.Errorf("modular inverse does not exist for denominator")
		}

		lagrangeCoeff := new(big.Int).Mul(numerator, denominatorInverse)
		lagrangeCoeff.Mod(lagrangeCoeff, modulus)

		term := new(big.Int).Mul(share.Y, lagrangeCoeff)
		term.Mod(term, modulus)

		secret.Add(secret, term)
		secret.Mod(secret, modulus)
	}

	return secret, nil
}
