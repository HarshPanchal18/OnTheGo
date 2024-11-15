package main

import (
	"hash/crc32"
	"os"
)

func hashedBytes(bytes []byte) uint32 {
	h := crc32.NewIEEE()
	h.Write(bytes)
	v := h.Sum32()
	return v
}

func getHash(filename string) (uint32, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil { return 0, err }
	return hashedBytes(bytes), nil
}
