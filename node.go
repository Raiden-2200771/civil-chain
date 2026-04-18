package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var bc = newBlockchain()

func chainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc.Blocks)
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

	w.WriteHeader(http.StatusCreated)
}
