package handler

import (
	// "fmt"
	"dasar/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome Sidiq, lets go belajar golang"))
	tmp, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error system, please wait!", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title" : "Golang Web",
	// 	"content" : "Learning golang for html",
	// }
	data := []entity.Product{
		{ ID: 1, Name: "Desle", Price: 259900, Stock: 12 },
		{ID: 2, Name: "Phoenix", Price: 229900, Stock: 5},
		{ID: 3, Name: "Motix", Price: 239900, Stock: 2},
	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error system, please wait!", http.StatusInternalServerError)
		return
	}
	// w.Write([]byte("Home"))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, lets go belajar golang"))
}

func MarioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mario from nintendo, lets go belajar golang"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}

	tmp, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error system, please wait!", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content" : idNum,
		"title" : "Product Web",
	}

	// data := entity.Product{ID: 1, Name: "Desle", Price: 229900, Stock: 12}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error system, please wait!", http.StatusInternalServerError)
		return
	}

	// w.Write([]byte("Product Page"))

	// fmt.Fprintf(w, "Product Page : %d", idNum)
}