package graph

// import (
// 	"context"
// 	"time"

// 	"github.com/ridwankustanto/shopvee/graphql/graph/model"
// )

// type queryResolver struct {
// 	server *Server
// }

// func (r *queryResolver) Accounts(ctx context.Context, pagination *model.PaginationInput, id *string) ([]*model.Account, error) {

// 	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
// 	defer cancel()

// 	if id != nil {
// 		a, err := r.server.accountClient.FindAccount(ctx, *id)
// 		if err != nil {
// 			return nil, err
// 		}

// 		return []*model.Account{{
// 			ID:   a.Id,
// 			Name: a.Name,
// 		}}, nil
// 	}

// 	accountList, err := r.server.accountClient.GetAccounts(ctx, uint64(*pagination.Skip), uint64(*pagination.Take))
// 	if err != nil {
// 		return nil, err
// 	}

// 	accounts := []*model.Account{}
// 	for _, a := range accountList {
// 		accounts = append(accounts, &model.Account{
// 			ID:   a.Id,
// 			Name: a.Name,
// 		})
// 	}

// 	return accounts, nil
// }
