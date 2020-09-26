package user

/*
import (
	"go-rest-sample/domain/entity/user"
	"go-rest-sample/domain/repository/user"
  "database/sql"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
    return &userPersistence{}
}

//ユーザ登録
func (up userPersistence) InsertUser(DB *sql.DB, userID, name, email string) error {
    stmt, err := DB.Prepare("INSERT INTO user(user_id, name, email) VALUES(?, ?, ?)")
    if err != nil {
        return err
    }
    _, err = stmt.Exec(userID, name, email)
    return err
}

//userIDによってユーザ情報を取得する
func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
    row := DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
    //row型をgolangで利用できる形にキャストする。
    return convertToUser(row)
}

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
