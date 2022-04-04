package order

import (
	context "context"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service OrderServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := NewOrderServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) CreateOrder(ctx context.Context, accountID string, products []*CreateOrderRequest_OrderProduct) (*Order, error) {

	// orderProduct := []*CreateOrderRequest_OrderProduct{}
	// for _, p := range products {
	// 	orderProduct = append(orderProduct, &CreateOrderRequest_OrderProduct{
	// 		ProductId: p.ProductId,
	// 		Quantity:  p.Quantity,
	// 	})
	// }
	orderRequest := &CreateOrderRequest{
		AccountId: accountID,
		Products:  products,
	}
	order, err := c.service.CreateOrder(ctx, orderRequest)
	if err != nil {
		return nil, err
	}

	return order.Order, nil
}

func (c *Client) GetOrderByAccountID(ctx context.Context, accountID string) ([]*Order, error) {
	orders, err := c.service.GetOrderByAccountID(ctx, &GetOrderByAccountIDRequest{AccountId: accountID})
	if err != nil {
		return nil, err
	}
	return orders.Order, nil
}
