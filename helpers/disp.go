package helpers

import (
	"encoding/hex"
	"fmt"

	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func DisplayDevice(device *device.Device) {
	// if device.PublicKey != nil {
	// 	fmt.Println("Public Key:")
	// 	fmt.Println("  X:", device.PublicKey.X)
	// 	fmt.Println("  Y:", device.PublicKey.Y)
	// } else {
	// 	fmt.Println("Public Key: nil")
	// }

	// if device.PrivateKey != nil {
	// 	fmt.Println("Private Key:", device.PrivateKey)
	// } else {
	// 	fmt.Println("Private Key: nil")
	// }

	fmt.Println("Share:")
	fmt.Println("  X:", device.Share.X)
	fmt.Println("  Y:", device.Share.Y)

	if len(device.SharedSecret) > 0 {
		fmt.Println("Shared Secret:")
		for i, secret := range device.SharedSecret {
			valx, valy, err := BytesToPair(secret)
			if err != nil {
				fmt.Println("Error converting secret to pair:", err)
				continue
			}
			fmt.Printf("  %d: X: %d Y: %d\n", i, valx, valy)
		}
	} else {
		fmt.Println("Shared Secret: nil")
	}

	if len(device.Chachakey) > 0 {
		fmt.Println("ChaCha Key:", hex.EncodeToString(device.Chachakey))
	} else {
		fmt.Println("ChaCha Key: nil")
	}

	if len(device.FinalHMAC) > 0 {
		fmt.Println("Final HMAC:", hex.EncodeToString(device.FinalHMAC))
	} else {
		fmt.Println("Final HMAC: nil")
	}

	fmt.Println("")
}

func DisplayDevices(devices []*device.Device) {
	for i, device := range devices {
		fmt.Printf("Device %d:\n", i+1)
		DisplayDevice(device)
	}
}

func Displayshares(shares []sss.Share) {
	fmt.Println("Shares:")
	for _, share := range shares {
		fmt.Printf("{X:%d Y:%d}\n", share.X, share.Y)
	}
	fmt.Println("")
}
