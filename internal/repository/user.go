package repository

import (
	"context"
	"fmt"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"

	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const users = "users"

type userRepository struct {
	db     *mongo.Database
	logger *logging.Logger
}

func NewUserRepository(db *mongo.Database, logger *logging.Logger) *userRepository {
	return &userRepository{db: db, logger: logger}
}

func (ur *userRepository) Create(ctx context.Context, user domain.User) error {
	ur.logger.Debug("create user")
	result, err := ur.db.Collection(users).InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user due to error: %v", err)
	}

	ur.logger.Debug("convert insertedID to objectID")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return nil
	}
	ur.logger.Trace(user)
	return fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
}

func (ur *userRepository) FindByID(ctx context.Context, id string) (u domain.User, err error) {
	ur.logger.Debug("convert id to ObjectID format")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert Hex to objextid. hex: %s", id)
	}

	filter := bson.M{"_id": oid}

	ur.logger.Debugf("find user with id: %s", id)
	result := ur.db.Collection(users).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperrors.ErrUserNotFound
		}
		return u, fmt.Errorf("failed to find user by id: %s due to error: %v", id, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user id: %s drom db due to error: %v", id, err)
	}
	ur.logger.Trace(u)

	return u, nil
}

func (ur *userRepository) FindByEmail(ctx context.Context, email string) (u domain.User, err error) {
	ur.logger.Debugf("find user with email: %s", email)
	filter := bson.M{"email": email}

	result := ur.db.Collection(users).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperrors.ErrUserNotFound
		}
		return u, fmt.Errorf("failed to find user by email: %s due to error: %v", email, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user with email: %s drom db due to error: %v", email, err)
	}
	ur.logger.Trace(u)

	return u, nil
}

func (ur *userRepository) FindByUsername(ctx context.Context, username string) (u domain.User, err error) {
	ur.logger.Debugf("find user with username: %s", username)
	filter := bson.M{"username": username}

	result := ur.db.Collection(users).FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return u, apperrors.ErrUserNotFound
		}
		return u, fmt.Errorf("failed to find user by username: %s due to error: %v", username, err)
	}

	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode user with username: %s drom db due to error: %v", username, err)
	}
	ur.logger.Trace(u)

	return u, nil
}

func (ur *userRepository) FindAll(ctx context.Context) (u []domain.User, err error) {
	ur.logger.Debug("find all users")
	cursor, err := ur.db.Collection(users).Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to find all users due to error: %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("failed to read all documents from cursor: %v", err)
	}

	return u, nil
}
func (ur *userRepository) Update(ctx context.Context, userDTO domain.UpdateUserDTO, id string) error {
	ur.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(userDTO)
	if err != nil {
		return fmt.Errorf("failed to marshal user. error: %v", err)
	}

	var updateUserObj bson.M

	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal user bytes. error: %v", err)
	}

	if userDTO.Email == "" {
		delete(updateUserObj, "email")
	}
	if userDTO.Username == "" {
		delete(updateUserObj, "username")
	}

	update := bson.M{
		"$set": updateUserObj,
	}

	ur.logger.Debugf("update user with id: %s", id)
	result, err := ur.db.Collection(users).UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failde to execute update user query. error: %v", err)
	}

	if result.MatchedCount == 0 {
		return apperrors.ErrUserNotFound
	}
	ur.logger.Tracef("Matched %d documents and modifed %d", result.MatchedCount, result.ModifiedCount)

	return nil
}
func (ur *userRepository) Delete(ctx context.Context, id string) error {
	ur.logger.Debug("convert id to ObjectID format")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failed to convert user ID to objectID. ID = %s", id)
	}
	filter := bson.M{"_id": objectID}

	ur.logger.Debugf("delete user with id: %s", id)
	result, err := ur.db.Collection(users).DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to execute query. error: %v", err)
	}
	if result.DeletedCount == 0 {
		return apperrors.ErrUserNotFound
	}
	ur.logger.Tracef("Deleted %d documents", result.DeletedCount)

	return nil
}
