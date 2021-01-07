// Author: magician
// File:   demo1.go
// Date:   2020/5/5
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
