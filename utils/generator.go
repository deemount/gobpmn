package utils

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
)

// GenerateHash ...
func GenerateHash() string {

	n := 8
	b := make([]byte, n)
	h := fnv.New32a()
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x", b)

	if _, err := h.Write([]byte(s)); err != nil {
		panic(err)
	}
	defer h.Reset()

	return fmt.Sprintf("%x", h.Sum(nil))

}
