package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", findAllUsers)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Fix .
func findAllUsers(w http.ResponseWriter, r *http.Request) {

	// Connect to DB
	db, err := gorm.Open("mysql", "root:root@/sample?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatalf("DB connection failed %v .", err)
	}
	// Display detailed logs .
	db.LogMode(true)

	// Empty slice .
	var userList []User
	db.Table("users").Find(&userList)

	response, _ := json.Marshal(userList)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
