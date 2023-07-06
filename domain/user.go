package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	Uuid            string             `bson:"uuid"`
	Email           string             `bson:"email"`
	UserName        string             `bson:"username"`
	DisplayName     string             `bson:"display_name"`
	Avatar          string             `bson:"avatar"`
	PhoneNumber     string             `bson:"phone_number"`
	AccountStatus   string             `bson:"account_status"`
	IsActive        bool               `bson:"is_active"`
	IsTermsAccepted bool               `bson:"is_terms_accepted"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) error
	ValidateUser(ctx context.Context, email string, uuid string) (User, error)
	FindAndUpdateUser(ctx context.Context, userId string, userData *User) error
	FindUser(ctx context.Context, userId string) (User, error)
}
