package sss

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project_support/models"
)

func ReconstructPolynomial(shares []*models.Share) (*big.Int, error) {
	if len(shares) < 2 {
		return nil, fmt.Errorf("at least 2 shares are required for reconstruction")
	}

	secret := big.NewInt(0)
	modulus := big.NewInt(10000000999)

	for i, share := range shares {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)

		for j, otherShare := range shares {
			if i != j {
				xDiff := new(big.Int).Sub(share.X, otherShare.X)
				if xDiff.Cmp(big.NewInt(0)) == 0 {
					return nil, fmt.Errorf("duplicate share X values detected")
				}

				numerator.Mul(numerator, new(big.Int).Neg(otherShare.X))
				numerator.Mod(numerator, modulus)

				denominator.Mul(denominator, xDiff)
				denominator.Mod(denominator, modulus)
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
