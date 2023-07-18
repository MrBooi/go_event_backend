package repository

import (
	"context"

	"github.com/mrbooi/event_backend/domain"
	"github.com/mrbooi/event_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (ur *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)

	_, err := collection.InsertOne(ctx, user)

	return err
}

// FindAndUpdateUser implements domain.UserRepository.
func (ur *userRepository) FindAndUpdateUser(ctx context.Context, userId string, userData *domain.User) error {
	panic("unimplemented")
}

// FindUser implements domain.UserRepository.
func (ur *userRepository) FindUser(ctx context.Context, userId string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return user, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": idHex}).Decode(&user)
	return user, err

}

// ValidateUser implements domain.UserRepository.
func (ur *userRepository) ValidateUser(ctx context.Context, email string, uuid string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User

	err := collection.FindOne(ctx, bson.M{"email": email, "uuid": uuid}).Decode(&user)
	return user, err
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
