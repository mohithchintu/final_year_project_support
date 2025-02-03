package main

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/hmac"
	"github.com/mohithchintu/final_year_project_support/models"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func DisplayDevices(devices []*models.Device) {
	for _, device := range devices {
		fmt.Printf("Device ID: %s\n", device.ID)
		fmt.Printf("  Private Key: %s\n", device.PrivateKey.String())
		fmt.Printf("  Threshold: %d\n", device.Threshold)
		fmt.Printf("  Group Key: %s\n", device.GroupKey.String())
		fmt.Printf("  Peers: ")
		for _, peer := range device.Peers {
			fmt.Printf("%s ", peer.ID)
		}
		for i, coeff := range device.Coefficients {
			fmt.Printf("\n  Coefficient %d: %s", i, coeff.String())
		}
		for _, share := range device.Shares {
			fmt.Printf("\n  Share X: %s, Y: %s", share.X.String(), share.Y.String())
		}
		fmt.Println("\n---------------------")
	}
}

func main() {
	device1 := models.NewDevice("Device1", big.NewInt(123456789), 3)
	device2 := models.NewDevice("Device2", big.NewInt(987654321), 3)
	device3 := models.NewDevice("Device3", big.NewInt(135790864), 3)
	device4 := models.NewDevice("Device4", big.NewInt(245672340), 3)

	devices := []*models.Device{device1, device2, device3, device4}
	// DisplayDevices(devices)
	for _, device := range devices {
		sss.GeneratePolynomial(device)
	}
	// DisplayDevices(devices)
	for _, device := range devices {
		device.Shares = sss.GenerateShares(device, device.Threshold+1)
	}
	// DisplayDevices(devices)

	device1.AddPeer(device2)
	device1.AddPeer(device3)
	device1.AddPeer(device4)
	device2.AddPeer(device1)
	device2.AddPeer(device3)
	device2.AddPeer(device4)
	device3.AddPeer(device1)
	device3.AddPeer(device2)
	device3.AddPeer(device4)
	device4.AddPeer(device1)
	device4.AddPeer(device2)
	device4.AddPeer(device3)

	for _, device := range devices {
		err := helpers.GenerateGroupKey(device)
		if err != nil {
			fmt.Printf("Error generating group key for %s: %v\n", device.ID, err)
		}
	}

	// DisplayDevices(devices)
	for _, device := range devices {
		message := "Hello Devices"
		hmacValue := hmac.ComputeHMAC(message, device.GroupKey)
		fmt.Printf("HMAC of Group Key for %s: %s\n", device.ID, hmacValue)
	}

	// testing
	// fmt.Println("\nReconstructing secrets for devices:")
	// for _, device := range devices {
	// 	reconstructedSecret, err := sss.ReconstructPolynomial(device.Shares)
	// 	if err != nil {
	// 		fmt.Printf("Error reconstructing secret for %s: %v\n", device.ID, err)
	// 	} else {
	// 		fmt.Printf("Reconstructed Secret for %s: %s\n", device.ID, reconstructedSecret.String())
	// 		message := "Verification Message"
	// 		hmacValue := hmac.ComputeHMAC(message, reconstructedSecret)
	// 		fmt.Printf("HMAC of Reconstructed Secret for %s: %s\n", device.ID, hmacValue)
	// 	}
	// }
}
