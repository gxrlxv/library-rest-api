package genre

import (
	"context"
	"github.com/gxrlxv/library-rest-api/internal/domain/genre"
	"go.mongodb.org/mongo-driver/mongo"
)

type genreStorage struct {
	db *mongo.Database
}

func NewStorage(db *mongo.Database) *genreStorage {
	return &genreStorage{db: db}
}

func (gs *genreStorage) Create(ctx context.Context, genre genre.Genre) (string, error) {
	return "", nil
}
func (gs *genreStorage) GetOne(ctx context.Context, id string) (genre.Genre, error) {
	return genre.Genre{}, nil
}
func (gs *genreStorage) GetAll(ctx context.Context) ([]genre.Genre, error) {
	return []genre.Genre{}, nil
}
func (gs *genreStorage) Update(ctx context.Context, genre genre.Genre) error {
	return nil
}
func (gs *genreStorage) Delete(ctx context.Context, id string) error {
	return nil
}
