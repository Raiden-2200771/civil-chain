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

func TestIsValidHash_ValidHash_ReturnsTrue(t *testing.T) {
	data := "014d82fc6825b4b2ca343134e7ca6297773a5e8779f6f9df16d2d8985c4052e9"

	got := isValidHash(data)

	if !got {
		t.Errorf("got false, want true")
	}
}

func TestIsValidHash_InvalidChars_ReturnsFalse(t *testing.T) {
	data := "014d82gc6825b4b2cg343134e7cz6297773a5e8779f6f9gf16d2d8985c4052z9"

	got := isValidHash(data)

	if got {
		t.Errorf("got true, want false")
	}
}

func TestIsValidHash_WrongLength_ReturnsFalse(t *testing.T) {
	data := "014d82fc6825b4b2ca343134e7ca6297773a5e8779f6f9df16d2d8985c4052e91"

	got := isValidHash(data)

	if got {
		t.Errorf("got true, want false")
	}
}

func TestLongestChain_SameLength_ReturnsA(t *testing.T) {
	a := []Block{{Index: 0, Data: "chain-a"}, {Index: 1, Data: "chain-a"}}
	b := []Block{{Index: 0, Data: "chain-b"}, {Index: 1, Data: "chain-b"}}

	got := longestChain(a, b)

	if got[0].Data != "chain-a" {
		t.Errorf("got %v, want chain-a", got[0].Data)
	}
}

func TestLongestChain_AisLonger_ReturnsA(t *testing.T) {
	a := []Block{{Index: 0}, {Index: 1}, {Index: 2}}
	b := []Block{{Index: 0}, {Index: 1}}

	got := longestChain(a, b)

	if len(got) != len(a) {
		t.Errorf("got length %d, want %d", len(got), len(a))
	}
}

func TestLongestChain_BisLonger_ReturnsB(t *testing.T) {
	a := []Block{{Index: 0}, {Index: 1}}
	b := []Block{{Index: 0}, {Index: 1}, {Index: 2}}

	got := longestChain(a, b)

	if len(got) != len(b) {
		t.Errorf("got length %d, want %d", len(got), len(b))
	}
}

func TestBlockHandler_InvalidHash_Returns400(t *testing.T) {
	body := strings.NewReader("これは生データです")
	req := httptest.NewRequest("POST", "/block", body)
	w := httptest.NewRecorder()

	blockHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("ステータスコードが正しくありません: got %d, want 400", w.Code)
	}
}

func TestBlockHandler_AddsBlock(t *testing.T) {
	body := strings.NewReader("014d82fc6825b4b2ca343134e7ca6297773a5e8779f6f9df16d2d8985c4052e9")
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
