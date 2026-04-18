package main

import (
	"strings"
	"testing"
)

func TestHash(t *testing.T) {
	b := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		Data:      "テストデータ",
		PrevHash:  "",
	}

	result := hash(b)

	if result == "" {
		t.Error("ハッシュ値が空です")
	}
}

func TestHashConsistency(t *testing.T) {
	b := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		Data:      "テストデータ",
		PrevHash:  "",
	}

	result1 := hash(b)
	result2 := hash(b)

	if result1 != result2 {
		t.Errorf("同じ入力なのに異なるハッシュが返りました: %s != %s", result1, result2)
	}
}

func TestHashDifference(t *testing.T) {
	b1 := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		Data:      "発言A",
		PrevHash:  "",
	}
	b2 := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		Data:      "発言B",
		PrevHash:  "",
	}

	result1 := hash(b1)
	result2 := hash(b2)

	if result1 == result2 {
		t.Errorf("異なるデータなのに同じハッシュが返りました: %s == %s", result1, result2)
	}
}

func TestNewBlockIndex(t *testing.T) {
	prev := Block{Index: 0}

	next := newBlock(prev, "データ")

	if next.Index != 1 {
		t.Errorf("Indexが正しくありません: got %d, want 1", next.Index)
	}
}

func TestNewBlockData(t *testing.T) {
	prev := Block{Index: 0}

	next := newBlock(prev, "政治家の発言")

	if next.Data != "政治家の発言" {
		t.Errorf("Dataが正しくありません: got %s, want 政治家の発言", next.Data)
	}
}

func TestNewBlockPrevHash(t *testing.T) {
	prev := Block{Index: 0, Timestamp: "2026-04-05T10:00:00Z", Data: "元のデータ"}

	next := newBlock(prev, "次のデータ")

	want := hash(prev)
	if next.PrevHash != want {
		t.Errorf("PrevHashが正しくありません: got %s, want %s", next.PrevHash, want)
	}
}

func TestNewBlockHash(t *testing.T) {
	prev := Block{Index: 0, Timestamp: "2026-04-05T10:00:00Z", Data: "元のデータ"}

	next := newBlock(prev, "次のデータ")

	want := hash(next)
	if next.Hash != want {
		t.Errorf("Hashが正しくありません: got %s, want %s", next.Hash, want)
	}
}

func TestNewBlockTimestamp(t *testing.T) {
	prev := Block{Index: 0}

	next := newBlock(prev, "次のデータ")

	if next.Timestamp == "" {
		t.Error("Timestampが空です")
	}
}

func TestMine_HashPrefix(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: "2026-04-18T10:00:00Z",
		Data:      "マイニングテスト",
		PrevHash:  "",
	}
	difficulty := 4

	result := mine(b, difficulty)

	prefix := strings.Repeat("0", difficulty)
	if !strings.HasPrefix(result.Hash, prefix) {
		t.Errorf("Hashが難易度%dの条件を満たしていません: got %s", difficulty, result.Hash)
	}
}

func TestMine_HashIsValid(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: "2026-04-18T10:00:00Z",
		Data:      "マイニングテスト",
		PrevHash:  "",
	}

	result := mine(b, 4)

	want := hash(result)
	if result.Hash != want {
		t.Errorf("Hashがblock内容と一致しません: got %s, want %s", result.Hash, want)
	}
}

func TestMine_NonceIsSet(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: "2026-04-18T10:00:00Z",
		Data:      "マイニングテスト",
		PrevHash:  "",
		Nonce:     0,
	}

	result := mine(b, 4)

	if result.Nonce < 0 {
		t.Errorf("Nonceが不正な値です: got %d", result.Nonce)
	}
}
