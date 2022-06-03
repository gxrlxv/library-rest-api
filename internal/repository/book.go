package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const books = "books"

type bookRepository struct {
	db     *mongo.Database
	logger *logging.Logger
}

func NewBookRepository(db *mongo.Database, logger *logging.Logger) *bookRepository {
	return &bookRepository{db: db, logger: logger}
}

func (bs *bookRepository) Create(ctx context.Context, book domain.Book) error {
	bs.logger.Debug("create book")
	result, err := bs.db.Collection(books).InsertOne(ctx, book)
	if err != nil {
		return fmt.Errorf("failed to create book due to error: %v", err)
	}

	bs.logger.Debug("convert insertedID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return nil
	}
	bs.logger.Trace(book)
	return fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}
func (bs *bookRepository) FindByID(ctx context.Context, id string) (domain.Book, error) {
	return domain.Book{}, nil
}
func (bs *bookRepository) FindAll(ctx context.Context) ([]domain.Book, error) {
	return []domain.Book{}, nil
}
func (bs *bookRepository) Update(ctx context.Context, book domain.Book) error {
	return nil
}
func (bs *bookRepository) Delete(ctx context.Context, id string) error {
	return nil
}
