package sss

import (
	"math/big"
	"math/rand"
)

type Share struct {
	X int
	Y int
}

const prime = 1000000007

func GenerateShares(secret int, n, k int) []Share {
	coefficients := make([]int, k)
	coefficients[0] = secret % prime

	for i := 1; i < k; i++ {
		coefficients[i] = rand.Intn(prime-1) + 1
	}

	shares := make([]Share, n)
	usedXValues := make(map[int]bool)

	// printPolynomial(coefficients)

	for i := 0; i < n; i++ {
		var x int
		for {
			x = rand.Intn(prime-1) + 1
			if !usedXValues[x] {
				usedXValues[x] = true
				break
			}
		}

		y := evaluatePolynomial(coefficients, x)
		shares[i] = Share{X: x, Y: y}
	}

	return shares
}

func evaluatePolynomial(coefficients []int, x int) int {
	y := 0
	for i, coef := range coefficients {
		y = (y + coef*modularPow(x, i, prime)) % prime
	}
	return y
}

func modularPow(base, exp, mod int) int {
	result := 1
	base = base % mod

	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func modularInverse(a, mod int) int {
	return int(new(big.Int).ModInverse(big.NewInt(int64(a)), big.NewInt(int64(mod))).Int64())
}

func LagrangeInterpolation(shares []Share) int {
	result := 0

	for i := range shares {
		numerator, denominator := 1, 1

		for j := range shares {
			if i != j {
				numerator = (numerator * (-shares[j].X + prime)) % prime
				denominator = (denominator * (shares[i].X - shares[j].X + prime)) % prime
			}
		}

		denominator = modularInverse(denominator, prime)
		term := (shares[i].Y * numerator % prime) * denominator % prime
		result = (result + term) % prime
	}

	return (result + prime) % prime
}
