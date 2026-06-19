package hash

import (
	"crypto/sha256"
	"encoding/json"
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

func (h SHA256) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.value)
}

func IsSHA256String(s string) bool {
	if len(s) != 64 {
		return false
	}
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	return true
}

func (h *SHA256) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if !IsSHA256String(s) {
		return fmt.Errorf("invalid SHA256 string: %s", s)
	}
	h.value = s
	return nil
}
