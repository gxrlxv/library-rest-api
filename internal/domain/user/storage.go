package user

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, dto CreateUserDTO) (string, error)
	GetOne(ctx context.Context, id string) (User, error)
	GetAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
