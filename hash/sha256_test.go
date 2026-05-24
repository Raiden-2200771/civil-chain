package hash_test

import (
	"testing"

	"civil-chain/hash"
)

func TestNew_SameDataReturnsSameHash(t *testing.T) {
	h1 := hash.New("政治家の発言")
	h2 := hash.New("政治家の発言")

	if h1 != h2 {
		t.Errorf("同じデータなのに異なるハッシュが返りました: %v != %v", h1, h2)
	}
}

func TestNew_DifferentDataReturnsDifferentHash(t *testing.T) {
	h1 := hash.New("政治家の発言")
	h2 := hash.New("別の発言")

	if h1 == h2 {
		t.Errorf("異なるデータなのに同じハッシュが返りました: %v", h1)
	}
}
