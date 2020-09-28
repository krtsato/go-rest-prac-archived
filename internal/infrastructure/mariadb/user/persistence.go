package user

import (
	"database/sql"
	"go-rest-sample/domain/entity/user"
	"go-rest-sample/domain/repository/user"
)

type userPersistence struct{}

func NewUserPersistence() user.Repository {
	return &userPersistence{}
}

// ユーザ登録
func (up userPersistence) InsertUser(DB *sql.DB, userEntity *user.entity) error {
	stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	return err
}

// userID によってユーザ情報を取得する
func (up userPersistence) SelectByUserID(DB *sql.DB, userID int) (*user.Entity, error) {
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
