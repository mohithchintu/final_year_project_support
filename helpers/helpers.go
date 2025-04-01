package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func GenerateRandomSecret() int {
	secret := make([]byte, 4)
	rand.Read(secret)
	return int(binary.LittleEndian.Uint32(secret))
}

func GenerateRandomKey32() []byte {
	key := make([]byte, 32)
	rand.Read(key)
	return key
}

func PairToBytes(a, b int) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, int32(a))
	binary.Write(buf, binary.LittleEndian, int32(b))
	return buf.Bytes()
}

func BytesToPair(data []byte) (int, int, error) {
	if len(data) < 8 {
		return 0, 0, fmt.Errorf("data is too short to be a valid pair")
	}

	var a, b int32
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &a)
	if err != nil {
		return 0, 0, fmt.Errorf("error reading first integer: %v", err)
	}
	err = binary.Read(buf, binary.LittleEndian, &b)
	if err != nil {
		return 0, 0, fmt.Errorf("error reading second integer: %v", err)
	}
	return int(a), int(b), nil
}
