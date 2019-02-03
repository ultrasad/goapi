package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

//Router is router
func Router() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
}

//RouterMux is router with mux
func RouterMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", homeMux).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	return r
}
