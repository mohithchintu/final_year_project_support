package sss

import (
	"fmt"
)

func printPolynomial(coefficients []int) {
	fmt.Print("Polynomial: ")
	for i, coef := range coefficients {
		if i == 0 {
			fmt.Printf("%d", coef)
		} else if coef != 0 {
			fmt.Printf(" + %d*x^%d", coef, i)
		}
	}
	fmt.Println()
}

func TestSSS() {
	secret := 67813484
	n := 3
	k := 3

	shares := GenerateShares(secret, n, k)
	fmt.Println("Generated Shares:")
	for _, share := range shares {
		fmt.Printf("X: %d, Y: %d\n", share.X, share.Y)
	}

	reconstructedSecret := LagrangeInterpolation(shares[:k])
	fmt.Printf("Reconstructed Secret: %d\n", reconstructedSecret)

	if reconstructedSecret == secret {
		fmt.Println("Secret reconstruction successful!")
	} else {
		fmt.Println("Secret reconstruction failed!")
	}
}
