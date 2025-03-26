package sss

import (
	"math/big"

	"github.com/mohithchintu/final_year_project_support/models"
)

// GenerateShares generates shares for a device based on its polynomial
func GenerateShares(device *models.Device, n int) []*models.Share {
	coefficients := device.Coefficients
	shares := make([]*models.Share, n)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		y := EvaluatePolynomial(coefficients, x)
		shares[i-1] = &models.Share{X: x, Y: y}
	}
	return shares
}

// EvaluatePolynomial evaluates the polynomial at a given point x
func EvaluatePolynomial(coefficients []*big.Int, x *big.Int) *big.Int {
	result := big.NewInt(0)
	temp := big.NewInt(0)

	for i := len(coefficients) - 1; i >= 0; i-- {
		temp.Mul(result, x)
		result.Add(temp, coefficients[i])
	}
	return result
}
