package utils

import "crypto/sha256"

func SHA256(data string) string {
	hash := sha256.Sum256(BytesOfString(data))
	return HexOfBytes(hash[:])
}
