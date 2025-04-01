package ecc

import (
	"fmt"
)

func TestECC() {
	fmt.Println("=== ECDH Key Exchange ===")
	dev1Priv, _ := GenerateECDHKeyPair()
	dev1Pub := &dev1Priv.PublicKey

	dev2Priv, _ := GenerateECDHKeyPair()
	dev2Pub := &dev2Priv.PublicKey

	fmt.Printf("device 1 sends P_A = (%x, %x)\n", dev1Pub.X, dev1Pub.Y)
	fmt.Printf("device 2 sends P_B = (%x, %x)\n", dev2Pub.X, dev2Pub.Y)

	fmt.Printf("device 1 P_B*d_a = %x\n", ComputeSharedSecret(dev1Priv, dev2Pub))
	fmt.Printf("device 2 P_A*d_b = %x\n", ComputeSharedSecret(dev2Priv, dev1Pub))

}
