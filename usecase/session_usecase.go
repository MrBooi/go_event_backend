package usecase

import (
	"context"
	"time"

	"github.com/mrbooi/event_backend/domain"
)

type sessionUseCase struct {
	sessionRepository domain.SessionRepository
	contextTimeout    time.Duration
}

// CreateUserSession implements domain.SessionUseCase.
func (su *sessionUseCase) CreateUserSession(ctx context.Context, session *domain.Session) error {
	cxt, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.sessionRepository.CreateUserSession(cxt, session)
}

// DeleteSession implements domain.SessionUseCase.
func (su *sessionUseCase) DeleteSession(ctx context.Context, sessionID string) error {
	cxt, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.sessionRepository.DeleteSession(cxt, sessionID)
}

// GetUserSessions implements domain.SessionUseCase.
func (su *sessionUseCase) GetUserSessions(ctx context.Context, uuid string) ([]domain.Session, error) {
	cxt, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.sessionRepository.GetUserSessions(cxt, uuid)
}

func NewSessionUseCase(sessionRepository domain.SessionRepository, timeout time.Duration) domain.SessionUseCase {
	return &sessionUseCase{
		sessionRepository: sessionRepository,
		contextTimeout:    timeout,
	}
}
