package utils

import (
	"crypto/sha1"
)

func SHA256(data string) string {
	hash := sha1.Sum(BytesOfString(data))
	return HexOfBytes(hash[:])
}
