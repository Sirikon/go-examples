package main

import (
	"fmt"
	"net/http"
)

func getNameFromRequest(r *http.Request) string {
	qs := r.URL.Query()
	if len(qs["name"]) > 0 {
		return qs["name"][0]
	}
	return ""
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	name := getNameFromRequest(r)
	fmt.Fprintf(w, "Hello %s!", name)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
