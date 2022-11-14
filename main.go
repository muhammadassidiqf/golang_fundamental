package main

import (
	"log"
	"net/http"
	"dasar/handler"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("About Page"))
	}

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/mario", handler.MarioHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	})
	mux.HandleFunc("/product", handler.ProductHandler)
	

	log.Println("starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

