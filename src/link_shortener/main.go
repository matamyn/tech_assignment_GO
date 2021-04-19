package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":50050"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Работй чертов виндовс! %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
