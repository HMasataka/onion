package repository

import (
	"context"

	"github.com/HMasataka/sqlboiler/domain/models"
	"github.com/HMasataka/sqlboiler/domain/repository"
	gotx "github.com/knocknote/gotx/rdbms"
)

type UserRepository struct {
	rdbms gotx.ClientProvider
}

func NewUserRepository(rdbms gotx.ClientProvider) repository.UserRepository {
	return &UserRepository{
		rdbms: rdbms,
	}
}

func (r *UserRepository) Find(ctx context.Context, userID string) (*models.User, error) {
	sqlboilerClient := r.rdbms.CurrentClient(ctx)
	return models.FindUser(ctx, sqlboilerClient, userID)
}
