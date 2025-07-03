package hw1

// go test ./hw1

func ToLittleEndian(n uint32) uint32 {
	return n&0xFF<<24 |
		n&0xFF00<<8 |
		n&0xFF0000>>8 |
		n&0xFF000000>>24
}
