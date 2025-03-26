package models

import (
	"math/big"
)

// Device represents an IoT device with its properties and associated peers.
type Device struct {
	ID           string
	PrivateKey   *big.Int
	PublicKey    *big.Int
	Shares       []*Share
	Peers        map[string]*Device
	Threshold    int
	GroupKey     *big.Int
	Coefficients []*big.Int
}

// NewDevice initializes a new device with the given ID, private key, and threshold for Shamir's Secret Sharing
func NewDevice(id string, privateKey *big.Int, threshold int) *Device {
	return &Device{
		ID:         id,
		PrivateKey: privateKey,
		Peers:      make(map[string]*Device),
		Shares:     make([]*Share, 0),
		Threshold:  threshold,
	}
}

// AddPeer adds another device as a peer to the current device
func (d *Device) AddPeer(peer *Device) {
	d.Peers[peer.ID] = peer
}
