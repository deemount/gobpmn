package utils

import (
	"crypto/rand"
	"fmt"
	"hash"
	"hash/fnv"
)

var fnvHash hash.Hash32 = fnv.New32a()

// GenerateHash ...
func GenerateHash() string {

	n := 8
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	fnvHash.Write([]byte(s))
	defer fnvHash.Reset()
	return fmt.Sprintf("%x", fnvHash.Sum(nil))
}
