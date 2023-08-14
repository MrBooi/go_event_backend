package usecase

import (
	"context"
	"time"

	"github.com/mrbooi/event_backend/domain"
)

type userUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// CreateUser implements domain.UserUserCase.
func (us *userUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.CreateUser(cxt, user)
}

// FindAndUpdateUser implements domain.UserUserCase.
func (us *userUseCase) FindAndUpdateUser(ctx context.Context, userId string, userData *domain.User) error {
	panic("unimplemented")
}

// FindUser implements domain.UserUserCase.
func (us *userUseCase) FindUser(ctx context.Context, userId string) (domain.User, error) {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.FindUser(cxt, userId)
}

// ValidateUser implements domain.UserUserCase.
func (us *userUseCase) ValidateUser(ctx context.Context, email string, uuid string) (domain.User, error) {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.ValidateUser(cxt, email, uuid)
}

func NewSignUpUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUserCase {
	return &userUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
