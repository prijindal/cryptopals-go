package cryptopalsgo

import (
	"encoding/hex"
)

func Xor(byte1 []byte, byte2 []byte) []byte {
	length := len(byte1)
	if length < len(byte2) {
		length = len(byte2)
	}
	byte3 := make([]byte, length)
	for i := 0; i < length; i++ {
		byte3[i] = byte1[i] ^ byte2[i]
	}
	return byte3
}

func XorHex(str1 string, str2 string) (string, error) {
	byte1, err := hex.DecodeString(str1)
	if err != nil {
		return "", nil
	}
	byte2, err := hex.DecodeString(str2)
	if err != nil {
		return "", nil
	}
	byte3 := Xor(byte1, byte2)
	return hex.EncodeToString(byte3), nil
}
