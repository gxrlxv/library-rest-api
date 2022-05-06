package repository

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorRepository struct {
	db *mongo.Database
}

func NewAuthorRepository(db *mongo.Database) *authorRepository {
	return &authorRepository{db: db}
}

func (as *authorRepository) Create(ctx context.Context, author domain.Author) (string, error) {
	return "", nil
}
func (as *authorRepository) GetOne(ctx context.Context, id string) (domain.Author, error) {
	return domain.Author{}, nil
}
func (as *authorRepository) GetAll(ctx context.Context) ([]domain.Author, error) {
	return []domain.Author{}, nil
}
func (as *authorRepository) Update(ctx context.Context, author domain.Author) error {
	return nil
}
func (as *authorRepository) Delete(ctx context.Context, id string) error {
	return nil
}
