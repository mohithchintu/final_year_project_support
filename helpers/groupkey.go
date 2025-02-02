package helpers

import (
	"fmt"
	"math/big"

	"github.com/mohithchintu/final_year_project_support/models"
	"github.com/mohithchintu/final_year_project_support/sss"
)

func GenerateGroupKey(device *models.Device) error {
	secret := big.NewInt(0)
	for _, peer := range device.Peers {
		part_secret, err := sss.ReconstructPolynomial(peer.Shares)
		if err != nil {
			fmt.Println("Error in reconstructing polynomial:", err)
			return err
		}
		secret.Add(secret, part_secret)
	}
	secret.Add(secret, device.PrivateKey)
	secret.Mod(secret, big.NewInt(10000000999))
	device.GroupKey = secret
	return nil
}
