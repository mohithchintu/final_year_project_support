package utils

import (
	"github.com/mohithchintu/final_year_project_support/chacha"
	"github.com/mohithchintu/final_year_project_support/device"
	"github.com/mohithchintu/final_year_project_support/ecc"
	"github.com/mohithchintu/final_year_project_support/helpers"
)

func DeviceConnect(device1 *device.Device, device2 *device.Device) {
	device1.Chachakey = ecc.ComputeSharedSecret(device1.PrivateKey, device2.PublicKey)
	device2.Chachakey = ecc.ComputeSharedSecret(device2.PrivateKey, device1.PublicKey)

	data1 := helpers.PairToBytes(device1.Share.X, device1.Share.Y)
	data2 := helpers.PairToBytes(device2.Share.X, device2.Share.Y)

	encryptKey1, nonce1 := chacha.EncryptChaCha20(data1, device1.Chachakey)
	encryptKey2, nonce2 := chacha.EncryptChaCha20(data2, device2.Chachakey)

	decryptKey1, err := chacha.DecryptChaCha20(encryptKey2, device1.Chachakey, nonce2)
	if err != nil {
		println("Error decrypting data1: ", err.Error())
		return
	}
	decryptKey2, err := chacha.DecryptChaCha20(encryptKey1, device2.Chachakey, nonce1)
	if err != nil {
		println("Error decrypting data1: ", err.Error())
		return
	}

	device1.SharedSecret = append(device1.SharedSecret, decryptKey1)
	device2.SharedSecret = append(device2.SharedSecret, decryptKey2)

	device1.Chachakey = nil
	device2.Chachakey = nil
}
