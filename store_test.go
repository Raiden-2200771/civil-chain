package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestSaveChain_FileCreated(t *testing.T) {
	bc := newBlockchain()
	path := "test_chain.json"
	defer os.Remove(path)

	err := saveChain(bc, path)

	if err != nil {
		t.Fatalf("saveChain() がエラーを返しました: %v", err)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("ファイルが作成されていません: %s", path)
	}
}

func TestSaveChain_ValidJSON(t *testing.T) {
	bc := newBlockchain()
	path := "test_chain.json"
	defer os.Remove(path)
	saveChain(bc, path)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ファイルの読み込みに失敗しました: %v", err)
	}
	if !json.Valid(data) {
		t.Errorf("ファイルの中身が有効なJSONではありません")
	}
}

func TestSaveChain_BlockCount(t *testing.T) {
	bc := newBlockchain()
	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")
	path := "test_chain.json"
	defer os.Remove(path)
	saveChain(bc, path)

	data, _ := os.ReadFile(path)
	var loaded Blockchain
	json.Unmarshal(data, &loaded)

	if len(loaded.Blocks) != len(bc.Blocks) {
		t.Errorf("ブロック数が一致しません: got %d, want %d", len(loaded.Blocks), len(bc.Blocks))
	}
}

func TestSaveChain_InvalidPath(t *testing.T) {
	bc := newBlockchain()

	err := saveChain(bc, "/invalid/path/chain.json")

	if err == nil {
		t.Errorf("エラーが返されるべきですが、nil でした")
	}
}

func TestChainFilePath_Port8080(t *testing.T) {
	got := chainFilePath("8080")
	want := "data/chain_8080.json"
	if got != want {
		t.Errorf("chainFilePath(\"8080\") = %q, want %q", got, want)
	}
}

func TestSaveChain_BlockData(t *testing.T) {
	bc := newBlockchain()
	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")
	path := "test_chain.json"
	defer os.Remove(path)
	saveChain(bc, path)

	data, _ := os.ReadFile(path)
	var loaded Blockchain
	json.Unmarshal(data, &loaded)

	for i, want := range bc.Blocks {
		got := loaded.Blocks[i]
		if got.Index != want.Index {
			t.Errorf("Block[%d].Index: got %d, want %d", i, got.Index, want.Index)
		}
		if got.Timestamp != want.Timestamp {
			t.Errorf("Block[%d].Timestamp: got %s, want %s", i, got.Timestamp, want.Timestamp)
		}
		if got.Data != want.Data {
			t.Errorf("Block[%d].Data: got %s, want %s", i, got.Data, want.Data)
		}
		if got.PrevHash != want.PrevHash {
			t.Errorf("Block[%d].PrevHash: got %s, want %s", i, got.PrevHash, want.PrevHash)
		}
		if got.Nonce != want.Nonce {
			t.Errorf("Block[%d].Nonce: got %d, want %d", i, got.Nonce, want.Nonce)
		}
		if got.Hash != want.Hash {
			t.Errorf("Block[%d].Hash: got %s, want %s", i, got.Hash, want.Hash)
		}
	}
}
