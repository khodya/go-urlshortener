package main

import (
	"log"
	"net/http"
	"net/url"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		w.WriteHeader(400)
	}
	if r.Method == "POST" {
		w.WriteHeader(201)
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
