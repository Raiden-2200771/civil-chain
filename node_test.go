package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChainHandler_ReturnsJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "/chain", nil)
	w := httptest.NewRecorder()

	chainHandler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ステータスコードが正しくありません: got %d, want 200", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Typeが正しくありません: got %s", w.Header().Get("Content-Type"))
	}
}

func TestFetchChain_Returns200_NoError(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[]`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	_, err := fetchChain(server.URL)

	if err != nil {
		t.Errorf("エラーは期待しないが、got: %v", err)
	}
}

func TestFetchChain_ReturnsNon200_Error(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	_, err := fetchChain(server.URL)

	if err == nil {
		t.Errorf("エラーを期待するが、got: nil")
	}
}

func TestFetchChain_ParsesBlocks(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"Index":1,"Timestamp":"2024-01-01","Data":"genesis","PrevHash":"abc","Nonce":42,"Hash":"xyz"}]`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	blocks, err := fetchChain(server.URL)

	if err != nil {
		t.Fatalf("エラーは期待しないが、got: %v", err)
	}
	if len(blocks) != 1 {
		t.Fatalf("ブロック数が正しくありません: got %d, want 1", len(blocks))
	}
	got := blocks[0]
	tests := []struct {
		name string
		got  any
		want any
	}{
		{"Index", got.Index, 1},
		{"Timestamp", got.Timestamp, "2024-01-01"},
		{"Data", got.Data, "genesis"},
		{"PrevHash", got.PrevHash, "abc"},
		{"Nonce", got.Nonce, 42},
		{"Hash", got.Hash, "xyz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func TestFetchChain_InvalidJSON_Error(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	_, err := fetchChain(server.URL)

	if err == nil {
		t.Errorf("エラーを期待するが、got: nil")
	}
}

func TestBlockHandler_AddsBlock(t *testing.T) {
	body := strings.NewReader("テスト発言")
	req := httptest.NewRequest("POST", "/block", body)
	w := httptest.NewRecorder()
	beforeCount := len(bc.Blocks)

	blockHandler(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("ステータスコードが正しくありません: got %d, want 201", w.Code)
	}
	if len(bc.Blocks) != beforeCount+1 {
		t.Errorf("ブロックが追加されていません: got %d, want %d", len(bc.Blocks), beforeCount+1)
	}
}
