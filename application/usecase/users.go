package usecase

import (
	"context"

	"github.com/HMasataka/onion/domain/models"
	"github.com/HMasataka/onion/domain/repository"
	"github.com/HMasataka/onion/transaction"
)

type UserUseCase interface {
	Find(ctx context.Context, userID string) (models.UserSlice, error)
}

type userUseCase struct {
	transactor     transaction.Transactor
	userRepository repository.UserRepository
}

func NewUserUseCase(transactor transaction.Transactor, userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		transactor:     transactor,
		userRepository: userRepository,
	}
}

func (u *userUseCase) Find(ctx context.Context, userID string) (models.UserSlice, error) {
	var user models.UserSlice
	err := u.transactor.Required(ctx, func(ctx context.Context) error {
		var err error
		user, err = u.userRepository.Find(ctx, userID)
		return err
	})
	return user, err
}
