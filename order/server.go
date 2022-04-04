package order

import (
	context "context"
	"fmt"
	"log"
	"net"

	"github.com/ridwankustanto/shopvee/account"
	"github.com/ridwankustanto/shopvee/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service       Service
	accountClient *account.Client
	productClient *product.Client
}

func ListenGRPC(s Service, accountUrl, productUrl string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()

	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		log.Fatal("failed on start new account client: ", err)
	}

	productClient, err := product.NewClient(productUrl)
	if err != nil {
		log.Fatal("failed on start new product client: ", err)
	}

	RegisterOrderServiceServer(serv, &grpcServer{
		service:       s,
		accountClient: accountClient,
		productClient: productClient,
	})

	reflection.Register(serv)
	return serv.Serve((lis))
}

func (s *grpcServer) CreateOrder(ctx context.Context, in *CreateOrderRequest) (*CreateOrderResponse, error) {
	// Check if account id exist
	_, err := s.accountClient.FindAccount(ctx, in.AccountId)
	if err != nil {
		log.Println("failed while check account:", err)
		return nil, err
	}

	// Check if product item exist
	products := []*Order_OrderProduct{}
	for _, orderProduct := range in.Products {
		p, err := s.productClient.FindProduct(ctx, orderProduct.ProductId)
		if err != nil {
			log.Println("failed while get data product:", err)
			return nil, err
		}
		products = append(products, &Order_OrderProduct{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    orderProduct.Quantity,
		})
	}

	// Create order
	order, err := s.service.CreateOrder(ctx, products, in.AccountId)
	if err != nil {
		log.Println("failed on create order:", err)
		return nil, err
	}

	return &CreateOrderResponse{
		Order: order,
	}, nil
}

func (s *grpcServer) GetOrderByAccountID(ctx context.Context, in *GetOrderByAccountIDRequest) (*GetOrderByAccountIDResponse, error) {
	// Check if account id exist
	_, err := s.accountClient.FindAccount(ctx, in.AccountId)
	if err != nil {
		log.Println("failed while check account:", err)
		return nil, err
	}

	// Get all order by account id
	listOrder, err := s.service.GetOrderByAccountID(ctx, in.AccountId)
	if err != nil {
		return nil, err
	}

	orders := []*Order{}
	for _, order := range listOrder {
		orders = append(orders, &Order{
			Id:         order.Id,
			AccountId:  order.AccountId,
			Timestamp:  order.Timestamp,
			TotalPrice: order.TotalPrice,
			Products:   order.Products,
		})
	}

	return &GetOrderByAccountIDResponse{
		Order: orders,
	}, nil
}

func (s *grpcServer) mustEmbedUnimplementedOrderServiceServer() {}
