package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/besta", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello."))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
