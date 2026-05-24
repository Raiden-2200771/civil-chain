package hash_test

import (
	"testing"

	"civil-chain/hash"
)

func TestNew_SameDataReturnsSameHash(t *testing.T) {
	h1 := hash.New("テスト公約")
	h2 := hash.New("テスト公約")

	if h1 != h2 {
		t.Errorf("同じデータなのに異なるハッシュが返りました: %v != %v", h1, h2)
	}
}

func TestNew_DifferentDataReturnsDifferentHash(t *testing.T) {
	h1 := hash.New("テスト公約")
	h2 := hash.New("別のテスト公約")

	if h1 == h2 {
		t.Errorf("異なるデータなのに同じハッシュが返りました: %v", h1)
	}
}

func TestZero_Returns64Zeros(t *testing.T) {
	h := hash.Zero()

	want := "0000000000000000000000000000000000000000000000000000000000000000"
	if h.String() != want {
		t.Errorf("Zero()が正しくありません: got %s, want %s", h.String(), want)
	}
}

func TestString_ReturnsSHA256Hash(t *testing.T) {
	h := hash.New("テスト公約")

	s := h.String()

	if len(s) != 64 {
		t.Errorf("String()の長さが不正です: got %d, want 64", len(s))
	}
	for _, c := range s {
		if !('0' <= c && c <= '9') && !('a' <= c && c <= 'f') {
			t.Errorf("String()に16進数以外の文字が含まれています: %c", c)
		}
	}
}
