package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
