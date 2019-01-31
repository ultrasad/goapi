package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type addressBook struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Code      int    `json:"code"`
	Phone     string `json:"phone"`
}

//getAddressBookAll is get address book all services
func getAddressBookAll(w http.ResponseWriter, r *http.Request) {

	fmt.Println("get address book all")

	addBook := addressBook{
		Firstname: "Chaiyarin",
		Lastname:  "Niamsuwan",
		Code:      1993,
		Phone:     "0870940955",
	}
	json.NewEncoder(w).Encode(addBook)
}

//HandleRequestAddress is call services method
func HandleRequestAddress() {
	fmt.Println("Handle Request Address")
	http.HandleFunc("/getAddress", getAddressBookAll)
	//http.ListenAndServe(":8080", nil)
}
