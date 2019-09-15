package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 1. Create router .
	router := mux.NewRouter().StrictSlash(true)
	// 2. Relate URLs to functions
	router.HandleFunc("/", home)
	// 3. Start on specified port
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
