package hash

import (
	"crypto/sha256"
	"fmt"
)

type SHA256 struct {
	value string
}

func New(data string) SHA256 {
	h := sha256.Sum256([]byte(data))
	return SHA256{value: fmt.Sprintf("%x", h)}
}

func Zero() SHA256 {
	return SHA256{value: "0000000000000000000000000000000000000000000000000000000000000000"}
}

func (h SHA256) String() string {
	return h.value
}
