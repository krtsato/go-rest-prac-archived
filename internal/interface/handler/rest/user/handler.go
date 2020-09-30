package user

import (
	"encoding/json"
	uuser "go-rest-sample/internal/usecase/user"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter" //Todo: gorilla/mux に変更
)

type Handler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type handler struct {
	useCase uuser.Usecase
}

// NewUserUseCase : User データに関する Handler を生成
func NewUserHandler(uu uuser.Usecase) Handler {
	return &handler{
		useCase: uu,
	}
}

// UserIndex : GET /users -> User データ一覧を返す
func (h handler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// request : 本 API のリクエストパラメータ
	//  -> こんな感じでリクエストも受け取れますが、今回は使いません
	type request struct {
		Begin uint `query:"begin"`
		Limit uint `query:"limit"`
	}

	// Field : response 内で使用する User を表す構造体
	//  -> ドメインモデルの User に HTTP の関心事である JSON タグを付与したくないために Handler 層で用意
	//     簡略化のために JSON タグを付与したドメインモデルを流用するプロジェクトもしばしば見かける

	type Field struct {
		ID        int        `json:"ID"`
		ProfileID int        `json:"profileID"`
		AuthID    string     `json:"authID"`
		Mail      string     `json:"mail"`
		Phone     string     `json:"phone"`
		CreatedAt *time.Time `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}

	// response : 本 API のレスポンス
	type response struct {
		Users []Field `json:"users"`
	}

	ctx := r.Context()

	// Usecase の呼出
	userSlice, err := h.useCase.SelectLimitedUsers(ctx, 2)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(response)
	for _, user := range *userSlice {
		field := Field(*user)
		res.Users = append(res.Users, field)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
