package main

import (
	"civil-chain/hash"
	"crypto/sha256"
	"fmt"
	"strings"
	"testing"
)

var dataHash = hash.Zero()

func TestHash(t *testing.T) {
	b := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		DataHash:  dataHash,
		PrevHash:  "",
	}

	result := hashBlock(b)

	if result == "" {
		t.Error("ハッシュ値が空です")
	}
}

func TestHashConsistency(t *testing.T) {
	b := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		DataHash:  dataHash,
		PrevHash:  "",
	}

	result1 := hashBlock(b)
	result2 := hashBlock(b)

	if result1 != result2 {
		t.Errorf("同じ入力なのに異なるハッシュが返りました: %s != %s", result1, result2)
	}
}

func TestHashDifference(t *testing.T) {
	b1 := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		DataHash:  dataHash,
		PrevHash:  "",
	}
	b2 := Block{
		Index:     0,
		Timestamp: "2026-04-04T10:00:00Z",
		DataHash:  hash.New("1"),
		PrevHash:  "",
	}

	result1 := hashBlock(b1)
	result2 := hashBlock(b2)

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

func TestNewBlockDataHash(t *testing.T) {
	prev := Block{Index: 0}

	next := newBlock(prev, "テスト公約")

	want := fmt.Sprintf("%x", sha256.Sum256([]byte("テスト公約")))
	if next.DataHash.String() != want {
		t.Errorf("DataHashが正しくありません: got %s, want %s", next.DataHash.String(), want)
	}
}

func TestNewBlockPrevHash(t *testing.T) {
	prev := Block{Index: 0, Timestamp: "2026-04-05T10:00:00Z", DataHash: dataHash}

	next := newBlock(prev, "次のデータ")

	want := hashBlock(prev)
	if next.PrevHash != want {
		t.Errorf("PrevHashが正しくありません: got %s, want %s", next.PrevHash, want)
	}
}

func TestNewBlockHash(t *testing.T) {
	prev := Block{Index: 0, Timestamp: "2026-04-05T10:00:00Z", DataHash: dataHash}

	next := newBlock(prev, "次のデータ")

	want := hashBlock(next)
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

func TestNewBlock_HashHasPrefix(t *testing.T) {
	prev := Block{Index: 0, Timestamp: "2026-04-18T10:00:00Z", DataHash: dataHash}

	next := newBlock(prev, "マイニング対応テスト")

	prefix := strings.Repeat("0", difficulty)
	if !strings.HasPrefix(next.Hash, prefix) {
		t.Errorf("HashがPoW条件を満たしていません: got %s", next.Hash)
	}
}

func TestMine_HashPrefix(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: "2026-04-18T10:00:00Z",
		DataHash:  dataHash,
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
		DataHash:  dataHash,
		PrevHash:  "",
	}

	result := mine(b, 4)

	want := hashBlock(result)
	if result.Hash != want {
		t.Errorf("Hashがblock内容と一致しません: got %s, want %s", result.Hash, want)
	}
}

func TestMine_NonceIsSet(t *testing.T) {
	b := Block{
		Index:     1,
		Timestamp: "2026-04-18T10:00:00Z",
		DataHash:  dataHash,
		PrevHash:  "",
		Nonce:     0,
	}

	result := mine(b, 4)

	if result.Nonce < 0 {
		t.Errorf("Nonceが不正な値です: got %d", result.Nonce)
	}
}
