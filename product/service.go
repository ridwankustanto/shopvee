package product

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error)
	FindProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip, take uint64) ([]Product, error)
}

type productService struct {
	respository Repository
}

func NewService(r Repository) Service {
	return &productService{respository: r}
}

func (s *productService) CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	p := Product{
		Id:          strings.ReplaceAll(uuid.NewString(), "-", ""),
		Name:        name,
		Description: description,
		Price:       price,
		Timestamp:   time.Now().Format("2006-01-02 15:04:05-0700"),
	}
	err := s.respository.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (s *productService) FindProduct(ctx context.Context, id string) (*Product, error) {
	p, err := s.respository.FindProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *productService) GetProducts(ctx context.Context, skip, take uint64) ([]Product, error) {
	res, err := s.respository.GetProducts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for _, p := range res {
		products = append(products, Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Timestamp:   p.Timestamp,
		})
	}
	return products, nil
}
