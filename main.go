package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	json2 "encoding/json"
)

type Item struct {
	Name string `json:"name"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var name string = r.Form.Get("name")
	var json, _ = json2.Marshal(&Item{Name: name})
	fmt.Fprint(w, string(json))
}