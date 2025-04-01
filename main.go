package main

import (
	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/sss"
	"github.com/mohithchintu/final_year_project_support/utils"
)

func main() {

	// hmac.TestHMAC()
	// chacha.TestChacha()
	// sss.TestSSS()
	// ecc.TestECC()

	numDevices := 7

	secret := helpers.GenerateRandomSecret()

	shares := sss.GenerateShares(secret, numDevices, numDevices)

	// helpers.Displayshares(shares)

	devices := make([]*device.Device, numDevices)

	for i, share := range shares {
		devices[i] = device.NewDevice(share, i+1)
	}

	// helpers.DisplayDevices(devices)

	for i, device := range devices {
		for j := i + 1; j < len(devices); j++ {
			otherDevice := devices[j]
			utils.DeviceConnect(device, otherDevice)
		}
	}

	// helpers.DisplayDevices(devices)

	for _, device := range devices {
		utils.ReconstructSecret(device)
	}

	helpers.DisplayDevices(devices)

}
