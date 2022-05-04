package author

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain/author"
	"go.mongodb.org/mongo-driver/mongo"
)

type authorStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) *authorStorage {
	return &authorStorage{db: db}
}

func (as *authorStorage) Create(ctx context.Context, author author.Author) (string, error) {
	return "", nil
}
func (as *authorStorage) GetOne(ctx context.Context, id string) (author.Author, error) {
	return author.Author{}, nil
}
func (as *authorStorage) GetAll(ctx context.Context) ([]author.Author, error) {
	return []author.Author{}, nil
}
func (as *authorStorage) Update(ctx context.Context, author author.Author) error {
	return nil
}
func (as *authorStorage) Delete(ctx context.Context, id string) error {
	return nil
}
