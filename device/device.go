package device

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/mohithchintu/final_year_project_support/ecc"
	"github.com/mohithchintu/final_year_project_support/sss"
)

type Peer struct {
	DeviceName string
	Share      []byte
}

type PeerHMAC struct {
	DeviceName string
	FinalHMAC  []byte
}

type Device struct {
	DeviceName  string
	PrivateKey  *ecdsa.PrivateKey
	PublicKey   *ecdsa.PublicKey
	SharedPeers []*Peer
	Share       sss.Share
	Chachakey   []byte
	HMACS       []*PeerHMAC
}

func NewDevice(share sss.Share, num int) *Device {
	privateKey, _ := ecc.GenerateECDHKeyPair()
	return &Device{
		DeviceName: fmt.Sprintf("Device%d", num),
		PrivateKey: privateKey,
		Share:      share,
		PublicKey:  &privateKey.PublicKey,
	}
}
