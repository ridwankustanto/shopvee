package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/ridwankustanto/shopvee/graphql/graph/generated"
	"github.com/ridwankustanto/shopvee/graphql/graph/model"
	"github.com/ridwankustanto/shopvee/order"
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	p, err := r.server.ProductClient.CreateProduct(ctx, product.Name, product.Description, product.Price)
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Timestamp:   p.Timestamp,
	}, nil
}
func (r *MutationResolver) CreateOrder(ctx context.Context, o model.OrderInput) (*model.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	products := []*order.CreateOrderRequest_OrderProduct{}
	for _, p := range o.Products {
		products = append(products, &order.CreateOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  uint32(p.Quantity),
		})
	}

	p, err := r.server.OrderClient.CreateOrder(ctx, o.AccountID, products)
	if err != nil {
		return nil, err
	}

	orderedProduct := []*model.OrderedProduct{}
	for _, p := range p.Products {
		orderedProduct = append(orderedProduct, &model.OrderedProduct{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    int(p.Quantity),
		})
	}
	return &model.Order{
		ID:         p.Id,
		TotalPrice: p.TotalPrice,
		Timestamp:  p.Timestamp,
		Products:   orderedProduct,
	}, nil
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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		p, err := r.server.ProductClient.FindProduct(ctx, *id)
		if err != nil {
			return nil, err
		}

		return []*model.Product{{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Timestamp:   p.Timestamp,
		}}, nil
	}

	productList, err := r.server.ProductClient.GetProducts(ctx, uint64(*pagination.Skip), uint64(*pagination.Take))
	if err != nil {
		return nil, err
	}

	products := []*model.Product{}
	for _, p := range productList {
		products = append(products, &model.Product{
			ID:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Timestamp:   p.Timestamp,
		})
	}

	return products, nil
}
func (r *QueryResolver) Orders(ctx context.Context, accountID *string) ([]*model.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	listOrder, err := r.server.OrderClient.GetOrderByAccountID(ctx, *accountID)
	if err != nil {
		return nil, err
	}

	orders := []*model.Order{}
	for _, o := range listOrder {
		orderProduct := []*model.OrderedProduct{}
		for _, p := range o.Products {
			orderProduct = append(orderProduct, &model.OrderedProduct{
				ID:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    int(p.Quantity),
			})
		}

		orders = append(orders, &model.Order{
			ID:         o.Id,
			TotalPrice: o.TotalPrice,
			Timestamp:  o.Timestamp,
			Products:   orderProduct,
		})
	}

	return orders, nil

}

func (r *Server) Mutation() generated.MutationResolver { return &MutationResolver{r} }
func (r *Server) Query() generated.QueryResolver       { return &QueryResolver{r} }

type MutationResolver struct {
	server *Server
}
type QueryResolver struct {
	server *Server
}
