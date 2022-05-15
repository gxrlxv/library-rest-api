package service

import (
	"context"
	"fmt"
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

func (us *userService) SignIn(ctx context.Context, user domain.User) error {
	model, err := us.repository.FindByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("user not found")
	}
	if model.Username != user.Username {
		return fmt.Errorf("wrong username")
	}
	if model.PasswordHash != user.PasswordHash {
		return fmt.Errorf("wrong password")
	}

	return nil
}

func (us *userService) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	return us.repository.FindByID(ctx, id)
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
