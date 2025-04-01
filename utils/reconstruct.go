package utils

import (
	"fmt"

	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/helpers"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func ReconstructSecret(device *device.Device) {
	shares := make([]sss.Share, 0)
	for _, share := range device.SharedSecret {
		valx, valy, err := helpers.BytesToPair(share)
		if err != nil {
			fmt.Println("Error converting secret to pair:", err)
			continue
		}
		shares = append(shares, sss.Share{X: valx, Y: valy})
	}
	shares = append(shares, device.Share)
	groupkey := sss.LagrangeInterpolation(shares)
	fmt.Println("Group Key: ", groupkey)
}
