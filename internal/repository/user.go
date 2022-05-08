package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db     *mongo.Database
	logger *logging.Logger
}

func NewUserRepository(db *mongo.Database, logger *logging.Logger) *userRepository {
	return &userRepository{db: db, logger: logger}
}

func (ur *userRepository) Create(ctx context.Context, user domain.User) (string, error) {
	ur.logger.Debug("create user")
	result, err := ur.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}

	ur.logger.Debug("convert insertedID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	ur.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}
func (ur *userRepository) GetOne(ctx context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}
func (ur *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
func (ur *userRepository) Update(ctx context.Context, user domain.User) error {
	return nil
}
func (ur *userRepository) Delete(ctx context.Context, id string) error {
	return nil
}
