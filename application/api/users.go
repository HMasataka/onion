package handler

import (
	"context"
	"net/http"

	"github.com/HMasataka/onion/application/api/router"
	"github.com/HMasataka/onion/application/usecase"
	"github.com/HMasataka/onion/domain/models"
)

type FindUserRequest struct {
	UserID string `json:"user_id"`
}

type FindUserResponse struct {
	User  *models.User   `json:"user_id"`
	Cards []*models.Card `json:"cards"`
}

type FindUserHandler struct {
	process func(ctx context.Context, req *FindUserRequest) (FindUserResponse, error)
}

func NewFindUserHandler(useCase usecase.UserUseCase) router.HandlerFunc {
	return &FindUserHandler{
		process: func(ctx context.Context, req *FindUserRequest) (FindUserResponse, error) {
			userSlice, err := useCase.Find(ctx, req.UserID)
			if err != nil {
				return FindUserResponse{}, err
			}
			// TODO UseCase層で詰め換え
			cards := make([]*models.Card, len(userSlice[0].R.UserCards))
			copy(cards, userSlice[0].R.UserCards)
			return FindUserResponse{User: userSlice[0], Cards: cards}, err
		},
	}
}

func (s *FindUserHandler) Process(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var req FindUserRequest
	if err := router.ReadRequest(r, &req); err != nil {
		return nil, err
	}
	return s.process(r.Context(), &req)
}
