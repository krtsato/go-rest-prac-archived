package user

import (
	"context"

	"go-rest-sample/domain/entity/user"
	"go-rest-sample/domain/repository/user"
)

type UserUsecase interface {
	SelectLimitedUsers(ctx context.Context, limit int) (*user.EntitySlice, error)
}

type userUsecase struct {
	// Repository や DomainService の手続き
	repository user.Repository
}

func NewUserUsecase(ur user.Repository) UserUsecase {
	return &userUsecase{
		repository: ur,
	}
}

func (uu userUsecase) SelectLimitedUsers(ctx context.Context) (user.EntitySlice, error) {
	// Persistence（Repository）を呼出
	userSlice, err := uu.repository.SelectLimitedUsers(ctx, 2)
	if err != nil {
		return nil, err
	}
	return userSlice, nil
}
