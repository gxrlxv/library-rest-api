package repository

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *userRepository {
	return &userRepository{db: db}
}

func (us *userRepository) Create(ctx context.Context, user domain.User) (string, error) {
	return "", nil
}
func (us *userRepository) GetOne(ctx context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}
func (us *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
func (us *userRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}
func (us *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}
