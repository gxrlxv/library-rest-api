package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (ur *userRepository) FindByID(ctx context.Context, id string) (u domain.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert Hex to objextid. hex: %s", id)
	}
	filter := bson.M{"_id": oid}

	result := ur.db.Collection("users").FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("not found")
		}
		return u, fmt.Errorf("failed to find user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user id: %s drom db due to error: %v", id, err)
	}

	return u, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (u domain.User, err error) {
	filter := bson.M{"email": email}

	result := ur.db.Collection("users").FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, fmt.Errorf("not found")
		}
		return u, fmt.Errorf("failed to find user by email: %s due to error: %v", email, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user with email: %s drom db due to error: %v", email, err)
	}

	return u, nil
}

func (ur *userRepository) FindAll(ctx context.Context) (u []domain.User, err error) {
	cursor, err := ur.db.Collection("users").Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to find all users due to error: %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("failed to read all documents from cursor: %v", err)
	}

	return u, nil
}
func (ur *userRepository) Update(ctx context.Context, user domain.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to objectID. ID = %s", user.ID)
	}
	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes. error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := ur.db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failde to execute update user query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}
	ur.logger.Tracef("Matched %d documents and modifed %d", result.MatchedCount, result.ModifiedCount)

	return nil
}
func (ur *userRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	result, err := ur.db.Collection("users").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failet to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}
	ur.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
