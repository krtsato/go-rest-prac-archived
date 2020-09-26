package userprofile

import (
	"database/sql"
	"go-rest-sample/internal/domain/entity/user"
)

type Repository interface {
	InsertUserprofile(DB *sql.DB, userEntity *user.Entity) (*user.Entity, error)
	SelectByUserprofileID(DB *sql.DB, userprofileID int) (*user.Entity, error)
}
