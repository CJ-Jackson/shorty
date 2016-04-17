package random

type RandomInterface interface {
	GenerateHex(numOfBytes int) string
}
