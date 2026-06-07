package hash_test

import (
	"encoding/json"
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

func TestNew_ReturnsValidSHA256(t *testing.T) {
	// value は非公開のため String() 経由でのみ検証できる
	h := hash.New("テスト公約")

	s := h.String()

	if len(s) != 64 {
		t.Errorf("New()の戻り値が64文字ではありません: got %d", len(s))
	}
	for _, c := range s {
		if !('0' <= c && c <= '9') && !('a' <= c && c <= 'f') {
			t.Errorf("New()の戻り値に16進数以外の文字が含まれています: %c", c)
		}
	}
}

// func TestString_ReturnsSHA256Hash は TestNew_ReturnsValidSHA256 で兼ねているため省略

func TestNew_DifferentDataReturnsDifferentHash(t *testing.T) {
	h1 := hash.New("テスト公約")
	h2 := hash.New("別のテスト公約")

	if h1 == h2 {
		t.Errorf("異なるデータなのに同じハッシュが返りました: %v", h1)
	}
}

func TestMarshalJSON_SerializesAsString(t *testing.T) {
	h := hash.New("テスト公約")

	b, err := json.Marshal(h)

	if err != nil {
		t.Fatalf("MarshalJSONでエラーが発生しました: %v", err)
	}
	want := `"` + h.String() + `"`
	if string(b) != want {
		t.Errorf("MarshalJSONが正しくありません: got %s, want %s", string(b), want)
	}
}

func TestUnmarshalJSON_DeserializesFromString(t *testing.T) {
	h := hash.New("テスト公約")
	b, _ := json.Marshal(h)

	var got hash.SHA256
	err := json.Unmarshal(b, &got)

	if err != nil {
		t.Fatalf("UnmarshalJSONでエラーが発生しました: %v", err)
	}
	if got != h {
		t.Errorf("UnmarshalJSONが正しくありません: got %s, want %s", got.String(), h.String())
	}
}

func TestZero_Returns64Zeros(t *testing.T) {
	h := hash.Zero()

	want := "0000000000000000000000000000000000000000000000000000000000000000"
	if h.String() != want {
		t.Errorf("Zero()が正しくありません: got %s, want %s", h.String(), want)
	}
}

