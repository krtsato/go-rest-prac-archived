package user

import (
	"database/sql"
	"go-rest-sample/internal/domain/entity/user"
)

type Repository interface {
	InsertUser(DB *sql.DB, userEntity *user.Entity) (*user.Entity, error)
	SelectByUserID(DB *sql.DB, userID int) (*user.Entity, error)
}
