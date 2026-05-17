package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	DataHash  string
	PrevHash  string
	Nonce     int
	Hash      string
}

func hash(b Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", b.Index, b.Timestamp, b.DataHash, b.PrevHash, b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func mine(b Block, difficulty int) Block {
	prefix := strings.Repeat("0", difficulty)

	for {
		b.Hash = hash(b)

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
		DataHash:  fmt.Sprintf("%x", sha256.Sum256([]byte(data))),
		PrevHash:  hash(prev),
	}
	b = mine(b, difficulty)
	return b
}
