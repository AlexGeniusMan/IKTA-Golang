package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home")
}

func contacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contacts")
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contacts)
	http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("Go рулит!")
	handleRequest()
}
