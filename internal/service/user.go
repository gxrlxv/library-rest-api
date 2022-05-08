package service

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
)

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *userService {
	return &userService{repository: repository}
}

func (us *userService) CreateUser(ctx context.Context, user domain.User) (string, error) {
	aid, err := us.repository.Create(ctx, user)
	if err != nil {
		panic(err)
	}
	return aid, err
}

func (us *userService) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return us.repository.GetOne(ctx, id)
}

func (us *userService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return us.repository.GetAll(ctx)
}
