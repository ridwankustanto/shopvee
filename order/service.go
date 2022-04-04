package order

import (
	context "context"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateOrder(ctx context.Context, products []*Order_OrderProduct, accountID string) (*Order, error)
	GetOrderByAccountID(ctx context.Context, accountID string) ([]Order, error)
}

type orderService struct {
	respository Repository
}

func NewService(r Repository) Service {
	return &orderService{r}
}

func (s *orderService) CreateOrder(ctx context.Context, products []*Order_OrderProduct, accountID string) (*Order, error) {
	var totalPrice float64
	for _, p := range products {
		price := float64(p.Quantity) * p.Price
		totalPrice += price
	}

	order := Order{
		Id:         strings.ReplaceAll(uuid.NewString(), "-", ""),
		AccountId:  accountID,
		Products:   products,
		TotalPrice: totalPrice,
		Timestamp:  time.Now().Format("2006-01-02 15:04:05-0700"),
	}
	err := s.respository.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (s *orderService) GetOrderByAccountID(ctx context.Context, accountID string) ([]Order, error) {
	return s.respository.GetOrderByAccountID(ctx, accountID)
}
