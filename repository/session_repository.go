package repository

import (
	"context"

	"github.com/mrbooi/event_backend/domain"
	"github.com/mrbooi/event_backend/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type sessionRepository struct {
	database   mongo.Database
	collection string
}

// CreateUserSession implements domain.SessionRepository.
func (sr *sessionRepository) CreateUserSession(ctx context.Context, session *domain.Session) error {
	collection := sr.database.Collection(sr.collection)

	_, err := collection.InsertOne(ctx, session)

	return err
}

// DeleteSession implements domain.SessionRepository.
func (sr *sessionRepository) DeleteSession(ctx context.Context, sessionID string) error {
	collection := sr.database.Collection(sr.collection)

	idHex, err := primitive.ObjectIDFromHex(sessionID)

	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(ctx, bson.D{{Key: "_id", Value: idHex}}, bson.D{{Key: "$set", Value: bson.D{{Key: "valid", Value: false}}}})

	return err
}

// GetUserSessions implements domain.SessionRepository.
func (sr *sessionRepository) GetUserSessions(ctx context.Context, uuid string) ([]domain.Session, error) {
	collection := sr.database.Collection(sr.collection)

	var sessions []domain.Session

	idHex, err := primitive.ObjectIDFromHex(uuid)

	if err != nil {
		return sessions, err
	}

	cursor, err := collection.Find(ctx, bson.M{"userId": idHex})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &sessions)

	if sessions == nil {

		return []domain.Session{}, err
	}

	return sessions, err
}

func NewSessionRepository(db mongo.Database, collection string) domain.SessionRepository {
	return &sessionRepository{
		database:   db,
		collection: collection,
	}
}
