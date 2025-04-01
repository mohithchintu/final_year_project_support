package hmac

import "fmt"

func TestHMAC() {
	data := "This is some test data"
	key := []byte("chintu_testing_val")

	generatedHMAC := GenerateHMAC(data, key)
	fmt.Printf("Generated HMAC: %x\n", generatedHMAC)

	if VerifyHMAC(data, key, generatedHMAC) {
		fmt.Println("HMAC verification succeeded")
	} else {
		fmt.Println("HMAC verification failed")
	}

	incorrectHMAC := append([]byte{}, generatedHMAC...)
	incorrectHMAC[0] ^= 0xFF

	fmt.Println("HMAC changed")
	if VerifyHMAC(data, key, incorrectHMAC) {
		fmt.Println("HMAC verification succeeded, but it should have failed")
	} else {
		fmt.Println("HMAC verification correctly failed")

	}

	incorrect_data := []byte("chintu")
	generatediHMAC := GenerateHMAC(data, incorrect_data)
	fmt.Printf("Generated incorrect HMAC: %x\n", generatediHMAC)
}
