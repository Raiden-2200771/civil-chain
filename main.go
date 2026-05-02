package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "ポート番号")
	flag.Parse()

	initNode(*port)

	http.HandleFunc("/chain", chainHandler)
	http.HandleFunc("/block", blockHandler)
	fmt.Println("Listening on :" + *port)
	http.ListenAndServe(":"+*port, nil)
}
