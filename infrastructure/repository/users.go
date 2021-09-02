package repository

import (
	"context"

	"github.com/HMasataka/sqlboiler/domain/models"
	"github.com/HMasataka/sqlboiler/domain/repository"
	"github.com/HMasataka/sqlboiler/transaction"
)

type UserRepository struct {
	connectionProvider transaction.ConnectionProvider
}

func NewUserRepository(connectionProvider transaction.ConnectionProvider) repository.UserRepository {
	return &UserRepository{
		connectionProvider: connectionProvider,
	}
}

func (r *UserRepository) Find(ctx context.Context, userID string) (*models.User, error) {
	client := r.connectionProvider.CurrentConnection(ctx)
	return models.FindUser(ctx, client, userID)
}
