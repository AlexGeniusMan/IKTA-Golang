package main

import (
	"fmt"
	"net/http"
)

type User struct {
	name      string
	age       uint8
	money     int32
	happiness float64
}

func (u *User) getName() string {
	return u.name
}

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
