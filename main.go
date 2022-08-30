package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome Sidiq, lets go belajar golang"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, lets go belajar golang"))
}

func marioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nintendo, lets go belajar golang"))
}
