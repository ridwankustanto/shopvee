package graph

// import (
// 	"context"
// 	"time"

// 	"github.com/ridwankustanto/shopvee/graphql/graph/model"
// )

// type mutationResolver struct {
// 	server *Server
// }

// func (r *mutationResolver) CreateAccount(ctx context.Context, account model.AccountInput) (*model.Account, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
// 	defer cancel()

// 	a, err := r.server.accountClient.CreateAccount(ctx, account.Name)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &model.Account{
// 		ID:   a.Id,
// 		Name: a.Name,
// 	}, nil
// }
