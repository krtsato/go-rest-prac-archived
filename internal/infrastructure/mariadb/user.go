package infrastructure

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
