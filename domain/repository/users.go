package repository

import (
	"context"

	"github.com/HMasataka/onion/domain/models"
)

type UserRepository interface {
	Find(ctx context.Context, userID string) (models.UserSlice, error)
}
