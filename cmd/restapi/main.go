package main

import (
	"fmt"
	"log"
	"net/http"

	puser "go-rest-sample/internal/infrastructure/mariadb/user"
	huser "go-rest-sample/internal/interfaces/handler/rest/user"
	uuser "go-rest-sample/internal/usecase/user"

	"github.com/julienschmidt/httprouter" // Todo: gorilla/mux に変更
)

func run() error {
	// Todo: Wire で DI
	userPersistence := puser.NewUserPersistence()
	userUseCase := uuser.NewUserUsecase(userPersistence)
	userHandler := huser.NewUserHandler(userUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/users", userHandler.Index)

	// サーバ起動
	fmt.Println("=====================================")
	fmt.Println("Server Start >> http://localhost:9999")
	fmt.Println("=====================================")
	log.Fatal(http.ListenAndServe(":9999", router))

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
