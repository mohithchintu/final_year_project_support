package chacha

import (
	"fmt"

	"github.com/mohithchintu/final_year_project_support/helpers"
)

func TestChacha() {
	a, b := 24, 1234
	fmt.Printf("Original pair: {%d, %d}\n", a, b)

	data := helpers.PairToBytes(a, b)
	fmt.Printf("Data as bytes: %x\n", data)

	key := helpers.GenerateRandomKey32()
	fmt.Printf("Generated random key: %x\n", key)

	encrypted, nonce := EncryptChaCha20(data, key)
	fmt.Printf("Encrypted data: %x\n", encrypted)
	fmt.Printf("Nonce: %x\n", nonce)

	decrypted, err := DecryptChaCha20(encrypted, key, nonce)
	if err != nil {
		fmt.Printf("Error during decryption: %v\n", err)
		return
	}
	fmt.Printf("Decrypted data: %x\n", decrypted)

	decryptedA, decryptedB, err := helpers.BytesToPair(decrypted)
	if err != nil {
		fmt.Printf("Error converting decrypted data back to pair: %v\n", err)
		return
	}
	fmt.Printf("Decrypted pair: {%d, %d}\n", decryptedA, decryptedB)

	if decryptedA == a && decryptedB == b {
		fmt.Println("Encryption and decryption are successful!")
	} else {
		fmt.Println("Decryption failed!")
	}
}
