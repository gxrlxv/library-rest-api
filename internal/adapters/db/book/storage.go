package book

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain/book"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) *bookStorage {
	return &bookStorage{db: db}
}

func (bs *bookStorage) Create(ctx context.Context, book book.Book) (string, error) {
	return "", nil
}
func (bs *bookStorage) GetOne(ctx context.Context, id string) (book.Book, error) {
	return book.Book{}, nil
}
func (bs *bookStorage) GetAll(ctx context.Context) ([]book.Book, error) {
	return []book.Book{}, nil
}
func (bs *bookStorage) Update(ctx context.Context, book book.Book) error {
	return nil
}
func (bs *bookStorage) Delete(ctx context.Context, id string) error {
	return nil
}
