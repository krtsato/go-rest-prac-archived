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
	row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
	//row型をgolangで利用できる形にキャストする。
	return convertToUser(row)
}

/*
//row型をuser型に紐づける
func convertToUser(row *sql.Row) (*domain.User, error) {
    user := domain.User{}
    err := row.Scan(&user.UserID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}
*/

/*
type UserDTO struct {
	Id   string
	Name string
}

func GetUserByID(conn *db.Conn, id int) (*UserDTO, error) { //DB にアクセスするロジック
	u := &UserDTO{}
	q := db.Query(`SELECT * FROM users WHERE id = "%d"`, id)
	if err := db.Exec(conn, q, u); err != nil {
		return nil, err
	}
	return u, nil
}
*/
