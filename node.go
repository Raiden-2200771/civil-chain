package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var bc = newBlockchain()
var chainPath string

func initNode(port string) {
	chainPath = chainFilePath(port)
	saveChain(bc, chainPath)
}

func chainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Blocks)
}

func fetchChain(url string) ([]Block, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ステータスコードが不正です: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var blocks []Block
	if err := json.Unmarshal(body, &blocks); err != nil {
		return nil, err
	}

	return blocks, nil
}

func longestChain(a, b []Block) []Block {
	if len(b) > len(a) {
		return b
	}
	return a
}

func blockHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからデータを読み込む
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "リクエストの読み込みに失敗しました", http.StatusBadRequest)
		return
	}

	// ブロックをチェーンに追加
	bc.addBlock(string(body))

	saveChain(bc, chainPath)

	w.WriteHeader(http.StatusCreated)
}
