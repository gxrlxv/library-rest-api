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

const authors = "authors"

type authorRepository struct {
	db     *mongo.Database
	logger *logging.Logger
}

func NewAuthorRepository(db *mongo.Database, logger *logging.Logger) *authorRepository {
	return &authorRepository{db: db, logger: logger}
}

func (ar *authorRepository) Create(ctx context.Context, author domain.Author) error {
	ar.logger.Debug("create author")
	result, err := ar.db.Collection(authors).InsertOne(ctx, author)
	if err != nil {
		return fmt.Errorf("failed to create author due to error: %v", err)
	}

	ar.logger.Debug("convert insertedID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return nil
	}
	ar.logger.Trace(author)
	return fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}
func (ar *authorRepository) FindByID(ctx context.Context, id string) (a domain.Author, err error) {
	ar.logger.Debug("convert id to ObjectID format")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return a, fmt.Errorf("failed to convert Hex to objextid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	ar.logger.Debugf("find author with id: %s", id)
	result := ar.db.Collection(authors).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return a, apperrors.ErrAuthorNotFound
		}
		return a, fmt.Errorf("failed to find author by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&a); err != nil {
		return a, fmt.Errorf("failed to decode author id: %s drom db due to error: %v", id, err)
	}
	ar.logger.Trace(a)

	return a, nil
}
func (ar *authorRepository) FindAll(ctx context.Context) (a []domain.Author, err error) {
	ar.logger.Debug("find all authors")
	cursor, err := ar.db.Collection(authors).Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return a, fmt.Errorf("failed to find all authors due to error: %v", err)
	}

	if err = cursor.All(ctx, &a); err != nil {
		return a, fmt.Errorf("failed to read all documents from cursor: %v", err)
	}

	return a, nil
}

func (ar *authorRepository) FindByName(ctx context.Context, name string) (a domain.Author, err error) {
	ar.logger.Debugf("find aythor with name: %s", name)
	filter := bson.M{"name": name}

	result := ar.db.Collection(authors).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return a, apperrors.ErrAuthorNotFound
		}
		return a, fmt.Errorf("failed to find author by name: %s due to error: %v", name, err)
	}

	if err = result.Decode(&a); err != nil {
		return a, fmt.Errorf("failed to decode author with name: %s drom db due to error: %v", name, err)
	}
	ar.logger.Trace(a)

	return a, nil
}

func (ar *authorRepository) Update(ctx context.Context, authorDTO domain.UpdateAuthorDTO, id string) error {
	ar.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert author ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	authorBytes, err := bson.Marshal(authorDTO)
	if err != nil {
		return fmt.Errorf("failed to marshal author. error: %v", err)
	}

	var updateAuthorObj bson.M

	err = bson.Unmarshal(authorBytes, &updateAuthorObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal author bytes. error: %v", err)
	}

	if authorDTO.Name == "" {
		delete(updateAuthorObj, "name")
	}
	if authorDTO.Nationality == "" {
		delete(updateAuthorObj, "nationality")
	}
	
	update := bson.M{
		"$set": updateAuthorObj,
	}

	ar.logger.Debugf("update author with id: %s", id)
	result, err := ar.db.Collection(authors).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failde to execute update author query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperrors.ErrAuthorNotFound
	}
	ar.logger.Tracef("Matched %d documents and modifed %d", result.MatchedCount, result.ModifiedCount)

	return nil
}

func (ar *authorRepository) Delete(ctx context.Context, id string) error {
	ar.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert author ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	ar.logger.Debugf("delete author with id: %s", id)
	result, err := ar.db.Collection(authors).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return apperrors.ErrAuthorNotFound
	}
	ar.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
