package main

import (
	"civil-chain/hash"
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	DataHash  hash.SHA256
	PrevHash  string
	Nonce     int
	Hash      string
}

func hashBlock(b Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", b.Index, b.Timestamp, b.DataHash.String(), b.PrevHash, b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func mine(b Block, difficulty int) Block {
	prefix := strings.Repeat("0", difficulty)

	for {
		b.Hash = hashBlock(b)

		if strings.HasPrefix(b.Hash, prefix) {
			return b
		}
		b.Nonce++
	}
}

const difficulty = 4

func newBlock(prev Block, data string) Block {
	b := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().Format(time.RFC3339),
		DataHash:  hash.New(data),
		PrevHash:  hashBlock(prev),
	}
	b = mine(b, difficulty)
	return b
}
