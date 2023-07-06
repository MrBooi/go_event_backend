package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSession = "sessions"
)

type Session struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId"`
	Valid     bool               `bson:"valid" json:"valid"`
	UserAgent string             `bson:"userAgent" json:"userAgent"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
	UpdatedAt string             `bson:"updateAt" json:"updateAt"`
}

type SessionRepository interface {
	CreateUserSession(ctx context.Context, session *Session) error
	GetUserSessions(ctx context.Context, uuid string) ([]Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

type SessionUseCase interface {
	CreateUserSession(ctx context.Context, session *Session) error
	GetUserSessions(ctx context.Context, uuid string) ([]Session, error)
	DeleteSession(ctx context.Context, sessionID string) error
}
