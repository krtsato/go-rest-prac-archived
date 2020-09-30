package user

import (
	"context"
	"go-rest-sample/internal/domain/entity/user"
)

type Repository interface {
	InsertUser(ctx context.Context, userEntity *user.Entity) (*user.Entity, error)
	SelectByUserID(ctx context.Context, userID int) (*user.Entity, error)
}
