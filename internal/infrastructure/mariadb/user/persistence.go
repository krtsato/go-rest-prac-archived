package user

import (
	"context"
	euser "go-rest-sample/internal/domain/entity/user"
	ruser "go-rest-sample/internal/domain/repository/user"
	"time"
)

type Persistence struct{}

func NewUserPersistence() ruser.Repository {
	return &Persistence{}
}

/*
// ユーザ登録
func (up userPersistence) InsertUser(ctx context.Context, userEntity *user.entity) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

// userID によってユーザ情報を取得する
func (up userPersistence) SelectByUserID(ctx context.Context, userID int) (*user.Entity, error) {
	row := DB.QueryRow("SELECT * FROM user WHERE id = ?", userID)
	// row 型を golang で利用できる形にキャストする
	return convertToUser(row)
}

// row 型を user 型に紐づける
func convertToUser(row *sql.Row) (*user.Entity, error) {
	user := user.Entity{}
	err := row.Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
*/

func (p Persistence) SelectLimitedUsers(ctx context.Context, limit int) (*euser.EntitySlice, error) {
	// 一旦モックデータを返却する
	user1 := euser.Entity{
		ID:        1,
		ProfileID: 1,
		Mail:      "hoge@email.com",
		Phone:     "00011112222",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
		DeletedAt: &time.Time{},
	}

	user2 := euser.Entity{
		ID:        2,
		ProfileID: 2,
		Mail:      "fuga@email.com",
		Phone:     "33344445555",
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
		DeletedAt: &time.Time{},
	}

	return &euser.EntitySlice{&user1, &user2}, nil
}
