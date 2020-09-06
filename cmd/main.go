package main

import (
	"fmt"
	"log"
	"net/http"

	"./utils"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User User構造体
type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users", findAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", findByID).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	// 追加部分 DELETEで受ける
	router.HandleFunc("/users", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func findAllUsers(w http.ResponseWriter, r *http.Request) {
	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	var userList []User
	db.Find(&userList)

	// 共通化した処理を使う
	utils.RespondWithJSON(w, http.StatusOK, userList)
}

func findByID(w http.ResponseWriter, r *http.Request) {

	id, err := utils.GetID(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid parameter")
	}

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	var user User
	db.Where("id = ?", id).Find(&user)

	// 共通化した処理を使う
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// 共通処理で構造体を取得する
	var user User
	msg := utils.GetStruct(r, &user)
	if msg != "" {
		utils.RespondWithError(w, http.StatusBadRequest, msg)
		return
	}

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	// DBにINSERTする
	db.Create(&user)

	utils.RespondWithJSON(w, http.StatusOK, user)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	// 共通処理で構造体を取得する
	var user User
	msg := utils.GetStruct(r, &user)
	if msg != "" {
		utils.RespondWithError(w, http.StatusBadRequest, msg)
		return
	}

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	// Update実行
	db.Save(&user)
	// gormはSaveメソッドで主キーの部分をUpdateしてくれる。また、存在しないキーだったらINSERTされる

	utils.RespondWithJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// 共通処理で構造体を取得する
	var user User
	msg := utils.GetStruct(r, &user)
	if msg != "" {
		utils.RespondWithError(w, http.StatusBadRequest, msg)
		return
	}

	// IDがセットされていない場合はエラーを返す
	if user.ID == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "ID is not set .")
		return
	}

	// DB接続
	db := utils.GetConnection()
	defer db.Close()

	db.Delete(&user)

	utils.RespondWithJSON(w, http.StatusOK, user)
}
