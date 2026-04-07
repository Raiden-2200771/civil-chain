package main

import "time"

type Blockchain struct {
	Blocks []Block
}

func newBlockchain() Blockchain {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "Genesis Block",
		PrevHash:  "",
	}
	genesis.Hash = hash(genesis)

	return Blockchain{Blocks: []Block{genesis}}
}

func (bc *Blockchain) addBlock(data string) {
	prev := bc.Blocks[len(bc.Blocks)-1]
	b := newBlock(prev, data)
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) isTampered() bool {
	for i := 0; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		if current.Hash != hash(current) {
			return true
		}
	}
	return false
}
