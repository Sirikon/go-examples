package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{name}", indexHandler)
	http.ListenAndServe(":8000", r)
}
