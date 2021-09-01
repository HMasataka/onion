package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/HMasataka/sqlboiler/application/api/router"
	"github.com/HMasataka/sqlboiler/application/usecase"
	"github.com/HMasataka/sqlboiler/domain/models"
)

type FindUserRequest struct {
	UserID string `json:"user_id"`
}

type FindUserHandler struct {
	process func(ctx context.Context, req *FindUserRequest) (*models.User, error)
}

func NewFindUserHandler(useCase usecase.UserUseCase) router.HandlerFunc {
	return &FindUserHandler{
		process: func(ctx context.Context, req *FindUserRequest) (*models.User, error) {
			return useCase.Find(ctx, req.UserID)
		},
	}
}

func (s *FindUserHandler) Process(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var req FindUserRequest

	requestBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.New("")
	}

	err = json.Unmarshal(requestBytes, &req)
	if err != nil {
		return nil, errors.New("")
	}

	return s.process(r.Context(), &req)
}
