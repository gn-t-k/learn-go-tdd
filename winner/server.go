package main

import (
	"fmt"
	"net/http"
)

// TODO: https://andmorefine.gitbook.io/learn-go-with-tests/build-an-application/http-server#nitesutowoku-1

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "20")
}
