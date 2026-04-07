package main

import "testing"

func TestNewBlockchain_HashGenesisBlock(t *testing.T) {
	bc := newBlockchain()

	if len(bc.Blocks) != 1 {
		t.Errorf("ブロック数が正しくありません: got %d, want 1", len(bc.Blocks))
	}
}

func TestNewBlockchain_GenesisIndex(t *testing.T) {
	bc := newBlockchain()

	if bc.Blocks[0].Index != 0 {
		t.Errorf("GenesisブロックのIndexが正しくありません: got %d, want 0", bc.Blocks[0].Index)
	}
}

func TestAddBlock_Length(t *testing.T) {
	bc := newBlockchain()

	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")

	if len(bc.Blocks) != 2 {
		t.Errorf("ブロック数が正しくありません: got %d, want 2", len(bc.Blocks))
	}
}

func TestAddBlock_PrevHash(t *testing.T) {
	bc := newBlockchain()

	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")

	want := bc.Blocks[0].Hash
	got := bc.Blocks[1].PrevHash
	if got != want {
		t.Errorf("PrevHashが正しくありません: got %s, want %s", got, want)
	}
}

func TestAddBlock_Data(t *testing.T) {
	bc := newBlockchain()

	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")

	want := "田中議員が〇〇法案に賛成票を投じた"
	got := bc.Blocks[1].Data
	if got != want {
		t.Errorf("Dataが正しくありません: got %s, want %s", got, want)
	}
}

func TestIsTampered_NotTampered(t *testing.T) {
	bc := newBlockchain()
	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")

	got := bc.isTampered()

	if got != false {
		t.Errorf("正常なチェーンはfalseを返すべきです: got %t", got)
	}
}

func TestIsTampered_TamperedData(t *testing.T) {
	bc := newBlockchain()
	bc.addBlock("田中議員が〇〇法案に賛成票を投じた")
	bc.Blocks[1].Data = "田中議員が〇〇法案に反対票を投じた"

	got := bc.isTampered()

	if got != true {
		t.Errorf("Data改ざんを検知できていません: got %t", got)
	}
}
