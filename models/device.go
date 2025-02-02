package models

import (
	"math/big"
)

type Device struct {
	ID         string
	PrivateKey *big.Int
	Shares     []*Share
	Peers      map[string]*Device
	Threshold  int
	GroupKey   *big.Int
}

func NewDevice(id string, privateKey *big.Int) *Device {
	return &Device{
		ID:         id,
		PrivateKey: privateKey,
		Peers:      make(map[string]*Device),
		Shares:     make([]*Share, 0),
	}
}

func (d *Device) AddPeer(peer *Device) {
	d.Peers[peer.ID] = peer
}
