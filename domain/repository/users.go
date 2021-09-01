package repository

import (
	"context"

	"github.com/HMasataka/sqlboiler/domain/models"
)

type UserRepository interface {
	Find(ctx context.Context, userID string) (*models.User, error)
}
