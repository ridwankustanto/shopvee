package product

import (
	context "context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service ProductServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("failed on dialing grpc server:", err)
		return nil, err
	}
	c := NewProductServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) CreateProduct(ctx context.Context, name, description string, price float64) (*Product, error) {
	p, err := c.service.CreateProduct(ctx, &CreateProductRequest{Name: name, Description: description, Price: price})
	if err != nil {
		return nil, err
	}
	return &Product{
		Id:          p.Product.Id,
		Name:        p.Product.Name,
		Description: p.Product.Description,
		Price:       p.Product.Price,
		Timestamp:   p.Product.Timestamp,
	}, nil
}

func (c *Client) FindProduct(ctx context.Context, id string) (*Product, error) {
	p, err := c.service.FindProduct(ctx, &FindProductRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Product{
		Id:          p.Product.Id,
		Name:        p.Product.Name,
		Description: p.Product.Description,
		Price:       p.Product.Price,
		Timestamp:   p.Product.Timestamp,
	}, nil
}

func (c *Client) GetProducts(ctx context.Context, skip, take uint64) ([]Product, error) {
	res, err := c.service.GetProducts(ctx, &GetProductsRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}
	products := []Product{}
	for _, p := range res.Products {
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
