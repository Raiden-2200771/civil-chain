package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/chain", chainHandler)
	http.HandleFunc("/block", blockHandler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
