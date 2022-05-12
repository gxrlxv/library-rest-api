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
	return us.repository.FindOne(ctx, id)
}

func (us *userService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return us.repository.FindAll(ctx)
}

func (us *userService) UpdateUser(ctx context.Context, user domain.User) error {
	return us.repository.Update(ctx, user)
}
func (us *userService) DeleteUser(ctx context.Context, id string) error {
	return us.repository.Delete(ctx, id)
}
