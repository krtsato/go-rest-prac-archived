package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	// Add new relation .
	router.HandleFunc("/users", findAllUsers)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	// Data to be converted to JSON
	var userList = []User{
		User{Id: 1, FirstName: "Taro", LastName: "Yamada"},
		User{Id: 2, FirstName: "Jiro", LastName: "Sato"},
	}

	response, _ := json.Marshal(userList)
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
