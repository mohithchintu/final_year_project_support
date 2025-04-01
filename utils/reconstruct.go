package utils

import (
	"fmt"

	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/hmac"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func ReconstructSecret(dev *device.Device) {
	shares := make([]sss.Share, 0)
	for _, peer := range dev.SharedPeers {
		valx, valy, err := helpers.BytesToPair(peer.Share)
		if err != nil {
			fmt.Println("Error converting secret to pair:", err)
			continue
		}
		shares = append(shares, sss.Share{X: valx, Y: valy})
	}
	shares = append(shares, dev.Share)
	groupkey := sss.LagrangeInterpolation(shares)
	groupkeyBytes := helpers.IntToBytes(groupkey)

	hmacValue := hmac.GenerateHMAC(dev.DeviceName, groupkeyBytes)
	dev.HMACS = append(dev.HMACS, &device.PeerHMAC{
		DeviceName: dev.DeviceName,
		FinalHMAC:  hmacValue,
	})

	for _, peer := range dev.SharedPeers {
		hmacValue := hmac.GenerateHMAC(peer.DeviceName, groupkeyBytes)
		dev.HMACS = append(dev.HMACS, &device.PeerHMAC{
			DeviceName: peer.DeviceName,
			FinalHMAC:  hmacValue,
		})
	}

	dev.SharedPeers = nil
}
