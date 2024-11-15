package main

import (
	"crypto/sha1"
	"hash/crc32"
	"os"
)

func hashedBytes(bytes []byte) uint32 {
	hash := crc32.NewIEEE()
	hash.Write(bytes)
	value := hash.Sum32()
	return value
}

func getHash(filename string) (uint32, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil { return 0, err }
	return hashedBytes(bytes), nil
}

func sha1Hash(bytes []byte) []byte {
	hash := sha1.New()
	hash.Write(bytes)
	bytes = hash.Sum([]byte{})
	return bytes
}