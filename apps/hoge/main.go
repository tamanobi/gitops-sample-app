package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Hoge!")
	})
	log.Fatal(http.ListenAndServe(":1333", nil))
}
