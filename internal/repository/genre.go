package repository

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type genreRepository struct {
	db *mongo.Database
}

func NewGenreRepository(db *mongo.Database) *genreRepository {
	return &genreRepository{db: db}
}

func (gs *genreRepository) Create(ctx context.Context, genre domain.Genre) (string, error) {
	return "", nil
}
func (gs *genreRepository) GetOne(ctx context.Context, id string) (domain.Genre, error) {
	return domain.Genre{}, nil
}
func (gs *genreRepository) GetAll(ctx context.Context) ([]domain.Genre, error) {
	return []domain.Genre{}, nil
}
func (gs *genreRepository) Update(ctx context.Context, genre domain.Genre) error {
	return nil
}
func (gs *genreRepository) Delete(ctx context.Context, id string) error {
	return nil
}
