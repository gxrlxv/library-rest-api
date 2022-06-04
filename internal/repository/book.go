package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"go.mongodb.org/mongo-driver/bson"
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
func (bs *bookRepository) FindByID(ctx context.Context, id string) (b domain.Book, err error) {
	bs.logger.Debug("convert id to ObjectID format")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return b, fmt.Errorf("failed to convert Hex to objextid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	bs.logger.Debugf("find book with id: %s", id)
	result := bs.db.Collection(books).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return b, apperrors.ErrBookNotFound
		}
		return b, fmt.Errorf("failed to find book by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&b); err != nil {
		return b, fmt.Errorf("failed to decode book id: %s drom db due to error: %v", id, err)
	}
	bs.logger.Trace(b)

	return b, nil
}
func (bs *bookRepository) FindAll(ctx context.Context) (b []domain.Book, err error) {
	bs.logger.Debug("find all books")
	cursor, err := bs.db.Collection(books).Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return b, fmt.Errorf("failed to find all books  due to error: %v", err)
	}

	if err = cursor.All(ctx, &b); err != nil {
		return b, fmt.Errorf("failed to read all documents from cursor: %v", err)
	}

	return b, nil
}

func (bs *bookRepository) Update(ctx context.Context, book domain.Book, id string) error {
	bs.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert book ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	bookBytes, err := bson.Marshal(book)
	if err != nil {
		return fmt.Errorf("failed to marshal book. error: %v", err)
	}

	var updateBookObj bson.M

	err = bson.Unmarshal(bookBytes, &updateBookObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal book bytes. error: %v", err)
	}

	delete(updateBookObj, "_id")

	update := bson.M{
		"$set": updateBookObj,
	}

	bs.logger.Debugf("update book with id: %s", id)
	result, err := bs.db.Collection(books).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failde to execute update book query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperrors.ErrBookNotFound
	}
	bs.logger.Tracef("Matched %d documents and modifed %d", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (bs *bookRepository) Delete(ctx context.Context, id string) error {
	bs.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert book ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	bs.logger.Debugf("delete book with id: %s", id)
	result, err := bs.db.Collection(books).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return apperrors.ErrBookNotFound
	}
	bs.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
