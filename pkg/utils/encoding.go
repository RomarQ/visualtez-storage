package utils

import "encoding/hex"

func StringOfBytes(bytes []byte) string {
	return string(bytes[:])
}

func BytesOfString(data string) []byte {
	return []byte(data)
}

func HexOfBytes(bytes []byte) string {
	return hex.EncodeToString(bytes)
}
