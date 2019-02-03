package handlers

import (
	"fmt"
	"net/http"
)

//Home Controller
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World, Net/Http")
}

func homeMux(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello World, Mux")
}
