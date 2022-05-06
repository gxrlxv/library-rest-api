package repository

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookRepository struct {
	db *mongo.Database
}

func NewBookRepository(db *mongo.Database) *bookRepository {
	return &bookRepository{db: db}
}

func (bs *bookRepository) Create(ctx context.Context, book domain.Book) (string, error) {
	return "", nil
}
func (bs *bookRepository) GetOne(ctx context.Context, id string) (domain.Book, error) {
	return domain.Book{}, nil
}
func (bs *bookRepository) GetAll(ctx context.Context) ([]domain.Book, error) {
	return []domain.Book{}, nil
}
func (bs *bookRepository) Update(ctx context.Context, book domain.Book) error {
	return nil
}
func (bs *bookRepository) Delete(ctx context.Context, id string) error {
	return nil
}
