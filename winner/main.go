package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}

// つぎここ
// https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/io#nitesutowoku-1
