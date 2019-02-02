package main

import (
	"fmt"
	"log"
	"net/http"
)

//Counter is count view
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func _main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.ListenAndServe("localhost:8888", nil)
}

func main() {
	log.Print("Starting the service")

	/*
		http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, "Hello World")
		})
	*/

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":1323s", nil))
}

/*
func handleThing(w http.ResponseWriter, r *http.Request) {}

//convert to type func
//type handleThing func(w http.ResponseWriter, r *http.Request)

func handleFunc() {
	http.HandleFunc("/", check(handleThing))
	//http.HandleFunc("/", handleThing) // compiles correctly
}

func check(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Before")
		h.ServeHTTP(w, r) // call original
		log.Println("After")
	})
}
*/
