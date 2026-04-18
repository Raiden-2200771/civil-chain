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
