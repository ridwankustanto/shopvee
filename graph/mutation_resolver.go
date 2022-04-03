package main

import (
	"context"
	"time"

	"github.com/ridwankustanto/shopvee/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type mutationResolver struct {
	server *Server
}

func (r *mutationResolver) CreateAccount(ctx context.Context, account model.AccountInput) (*model.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.accountClient.CreateAccount(ctx, account.Name)
	if err != nil {
		return nil, err
	}

	return &model.Account{
		ID:   a.Id,
		Name: a.Name,
	}, nil
}
