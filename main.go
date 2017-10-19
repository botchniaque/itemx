package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	name string
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name string = r.Form.Get("name")
	fmt.Printf("%+v\n", Item{name})
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}