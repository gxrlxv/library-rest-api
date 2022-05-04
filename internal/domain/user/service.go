package user

import "context"

type Service interface {
	CreateUser(ctx context.Context, dto CreateUserDTO) (string, error)
	GetUserByID(ctx context.Context, id string) (User, error)
	GetAllUsers(ctx context.Context) ([]User, error)
}

type service struct {
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, dto CreateUserDTO) (string, error) {
	return s.storage.Create(ctx, dto)
}

func (s *service) GetUserByID(ctx context.Context, id string) (User, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *service) GetAllUsers(ctx context.Context) ([]User, error) {
	return s.storage.GetAll(ctx)
}
