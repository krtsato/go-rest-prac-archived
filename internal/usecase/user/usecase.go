package user

import (
	"context"

	euser "go-rest-sample/internal/domain/entity/user"
	ruser "go-rest-sample/internal/domain/repository/user"
)

type Usecase interface {
	SelectLimitedUsers(ctx context.Context, limit int) (*euser.EntitySlice, error)
}

type usecase struct {
	// Repository や DomainService の手続き
	repository ruser.Repository
}

func NewUserUsecase(ur ruser.Repository) Usecase {
	return &usecase{
		repository: ur,
	}
}

func (u usecase) SelectLimitedUsers(ctx context.Context, limit int) (*euser.EntitySlice, error) {
	// Persistence（Repository）を呼出
	userSlice, err := u.repository.SelectLimitedUsers(ctx, 2)
	if err != nil {
		return nil, err
	}
	return userSlice, nil
}
