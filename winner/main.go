package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}

// 次ここ
// https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/json#nitesutowoku-1
