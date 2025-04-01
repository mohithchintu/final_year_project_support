package helpers

import (
	"encoding/hex"
	"fmt"

	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func DisplayDevice(device *device.Device) {
	fmt.Println("Device Name:", device.DeviceName)
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

	if len(device.SharedPeers) > 0 {
		fmt.Println("Shared Secret:")
		for _, peer := range device.SharedPeers {
			fmt.Printf(" %s\n", peer.DeviceName)
			valx, valy, err := BytesToPair(peer.Share)
			if err != nil {
				fmt.Println("Error converting secret to pair:", err)
				continue
			}
			fmt.Printf(" X: %d Y: %d\n", valx, valy)
		}
	} else {
		fmt.Println("Shared Secret: nil")
	}

	if len(device.Chachakey) > 0 {
		fmt.Println("ChaCha Key:", hex.EncodeToString(device.Chachakey))
	} else {
		fmt.Println("ChaCha Key: nil")
	}

	if len(device.HMACS) > 0 {
		fmt.Println("HMACs:")
		for _, hmac := range device.HMACS {
			fmt.Printf(" %s\n", hmac.DeviceName)
			fmt.Printf(" HMAC: %s\n", hex.EncodeToString(hmac.FinalHMAC))
		}
	} else {
		fmt.Println("HMACs: nil")
	}

	fmt.Println("")
}

func DisplayDevices(devices []*device.Device) {
	for _, device := range devices {
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
