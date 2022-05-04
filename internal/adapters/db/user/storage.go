package user

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type userStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) *userStorage {
	return &userStorage{db: db}
}

func (us *userStorage) Create(ctx context.Context, user user.User) (string, error) {
	return "", nil
}
func (us *userStorage) GetOne(ctx context.Context, id string) (user.User, error) {
	return user.User{}, nil
}
func (us *userStorage) GetAll(ctx context.Context) ([]user.User, error) {
	return []user.User{}, nil
}
func (us *userStorage) Update(ctx context.Context, user user.User) error {
	return nil
}
func (us *userStorage) Delete(ctx context.Context, id string) error {
	return nil
}
