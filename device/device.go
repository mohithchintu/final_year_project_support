package device

import (
	"crypto/ecdsa"

	"github.com/mohithchintu/final_year_project_support/ecc"
	"github.com/mohithchintu/final_year_project_support/sss"
)

type Device struct {
	PrivateKey   *ecdsa.PrivateKey
	PublicKey    *ecdsa.PublicKey
	SharedSecret [][]byte
	Share        sss.Share
	Chachakey    []byte
	FinalHMAC    []byte
}

func NewDevice(share sss.Share) *Device {
	privateKey, _ := ecc.GenerateECDHKeyPair()

	return &Device{
		PrivateKey: privateKey,
		Share:      share,
		PublicKey:  &privateKey.PublicKey,
	}
}
