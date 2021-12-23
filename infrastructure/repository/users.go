package repository

import (
	"context"

	"github.com/HMasataka/onion/domain/models"
	"github.com/HMasataka/onion/domain/repository"
	"github.com/HMasataka/onion/transaction"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserRepository struct {
	connectionProvider transaction.ConnectionProvider
}

func NewUserRepository(connectionProvider transaction.ConnectionProvider) repository.UserRepository {
	return &UserRepository{
		connectionProvider: connectionProvider,
	}
}

func (r *UserRepository) Find(ctx context.Context, userID string) (models.UserSlice, error) {
	client := r.connectionProvider.CurrentConnection(ctx)
	return models.Users(qm.Load("UserCards")).All(ctx, client)
}
