package sss

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project_support/models"
)

// GeneratePolynomial generates the polynomial coefficients for a device using its private key
func GeneratePolynomial(device *models.Device) error {
	degree := device.Threshold
	if device.PrivateKey == nil {
		return fmt.Errorf("device secret (PrivateKey) is not set")
	}

	coefficients := make([]*big.Int, degree+1)
	coefficients[0] = device.PrivateKey

	maxVal := big.NewInt(10000000999)

	for i := 1; i <= degree; i++ {
		coeff, err := rand.Int(rand.Reader, maxVal)
		if err != nil {
			return fmt.Errorf("failed to generate random coefficient: %v", err)
		}
		coefficients[i] = coeff
	}
	device.Coefficients = coefficients
	return nil
}
