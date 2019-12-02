package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main(){
	http.HandleFunc("/", handlers)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8800", nil))
}

func handlers(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path(%d) = %q\n", count, r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}


