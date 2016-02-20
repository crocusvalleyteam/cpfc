package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Productname string `json: "productname"`
}

var prod = new(Product)

var sampleproducts = map[string]Product{"1": Product{Productname: "sample"}}

func main() {

	log.Fatal(http.ListenAndServe("localhost:8080", prod.router()))

}

func (prod Product) router() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/list", findallproducts).Methods("GET")

	return r

}

func findallproducts(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)

	w.Header().Set("content-type", "application/json")

	allproducts, err := json.Marshal(sampleproducts)

	if err != nil {

		prod.Handlerros(w, err)
	}

	fmt.Fprintf(w, string(allproducts))

}

//todo move this code to helpers
func (a Product) Handlerros(w http.ResponseWriter, e error) {

	if w != nil {

		fmt.Fprintf(w, e.Error())
	} else {

		fmt.Println(e.Error())
	}

}
