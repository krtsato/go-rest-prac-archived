package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// RespondWithError エラー情報をJSONで返す
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

// RespondWithJSON JSONを返す部分を共通化
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetConnection DBとのコネクションを張る
func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("DB connection failed %v", err)
	}
	db.LogMode(true)

	return db
}

// GetID リクエストからIDを取得する
func GetID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

// GetStruct JSONリクエストを構造体にパースする。エラーの場合はレスポンスするメッセージを返す。
func GetStruct(r *http.Request, i interface{}) string {

	// リクエストボディ取得
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return "Invalid request"
	}

	// 読み込んだJSONを構造体に変換
	if err := json.Unmarshal(body, i); err != nil {
		return "JSON Unmarshaling failed ."
	}

	return ""
}
