package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/ridwankustanto/shopvee/graphql/graph/generated"
	"github.com/ridwankustanto/shopvee/graphql/graph/model"
)

func (r *MutationResolver) CreateAccount(ctx context.Context, account model.AccountInput) (*model.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	a, err := r.server.AccountClient.CreateAccount(ctx, account.Name)
	if err != nil {
		return nil, err
	}

	return &model.Account{
		ID:   a.Id,
		Name: a.Name,
	}, nil
}

func (r *MutationResolver) CreateProduct(ctx context.Context, product model.ProductInput) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *QueryResolver) Accounts(ctx context.Context, pagination *model.PaginationInput, id *string) ([]*model.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		a, err := r.server.AccountClient.FindAccount(ctx, *id)
		if err != nil {
			return nil, err
		}

		return []*model.Account{{
			ID:   a.Id,
			Name: a.Name,
		}}, nil
	}

	accountList, err := r.server.AccountClient.GetAccounts(ctx, uint64(*pagination.Skip), uint64(*pagination.Take))
	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}
	for _, a := range accountList {
		accounts = append(accounts, &model.Account{
			ID:   a.Id,
			Name: a.Name,
		})
	}

	return accounts, nil
}

func (r *QueryResolver) Products(ctx context.Context, pagination *model.PaginationInput, id *string) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Server) Mutation() generated.MutationResolver { return &MutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Server) Query() generated.QueryResolver { return &QueryResolver{r} }

type MutationResolver struct {
	server *Server
}
type QueryResolver struct {
	server *Server
}
