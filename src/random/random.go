package random

import (
	"crypto/rand"
	"fmt"
)

type Random struct{}

func (_ Random) GenerateHex(numOfBytes int) string {
	data := make([]byte, numOfBytes)
	rand.Read(data)

	return fmt.Sprintf("%x", data)
}
