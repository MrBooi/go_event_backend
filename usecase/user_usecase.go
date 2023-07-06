package usecase

import (
	"context"
	"time"

	"github.com/mrbooi/event_backend/domain"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// CreateUser implements domain.UserUserCase.
func (us *userUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.CreateUser(cxt, user)
}

// FindAndUpdateUser implements domain.UserUserCase.
func (us *userUsecase) FindAndUpdateUser(ctx context.Context, userId string, userData *domain.User) error {
	panic("unimplemented")
}

// FindUser implements domain.UserUserCase.
func (us *userUsecase) FindUser(ctx context.Context, userId string) (domain.User, error) {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.FindUser(cxt, userId)
}

// ValidateUser implements domain.UserUserCase.
func (us *userUsecase) ValidateUser(ctx context.Context, email string, uuid string) (domain.User, error) {
	cxt, cancel := context.WithTimeout(ctx, us.contextTimeout)
	defer cancel()
	return us.userRepository.ValidateUser(cxt, email,uuid)
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUserCase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
