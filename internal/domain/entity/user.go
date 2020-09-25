package entity

import "time"

type User struct {
	ID        int
	AuthID    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Profile   *Userprofile
}

func GetUserByID(conn *db.Conn, id int) (*User, error) {
	//インフラストラクチャレイヤの実装を利用する
	dto, err := mariadb.GetUserByID(conn, id)
	if err != nil {
		return nil, err
	}
	u := &User{
		Id:   dto.ID,
		Name: dto.Name,
	}
	return u, nil
}
