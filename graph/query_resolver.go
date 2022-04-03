package main

import (
	"context"
	"fmt"

	"github.com/ridwankustanto/shopvee/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *model.PaginationInput, id *string) ([]*model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}
