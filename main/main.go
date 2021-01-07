package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

//
//type User struct {
//	name      string
//	age       uint8
//	money     int32
//	happiness float64
//}
//
//func (u *User) getName() string {
//	return u.name
//}
//
//func home(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Home")
//}
//
//func contacts(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Contacts")
//}
//
//func handleRequest() {
//	http.HandleFunc("/", home)
//	http.HandleFunc("/contacts/", contacts)
//	http.ListenAndServe(":8000", nil)
//}

func main() {

	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // id последнего добавленного объекта
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	fmt.Println("Go рулит!")
	//handleRequest()
}
