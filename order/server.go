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
		p, err := s.getProductDetail(ctx, orderProduct.ProductId)
		if err != nil {
			log.Println("failed while get data product:", err)
			return nil, err
		}
		products = append(products, p)
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
		orderedProduct := []*Order_OrderProduct{}
		for _, p := range order.Products {
			product, err := s.getProductDetail(ctx, p.Id)
			if err != nil {
				log.Println("failed on get data product detail:", err)
				return nil, err
			}
			product.Quantity = p.Quantity

			orderedProduct = append(orderedProduct, product)
		}

		orders = append(orders, &Order{
			Id:         order.Id,
			AccountId:  order.AccountId,
			Timestamp:  order.Timestamp,
			TotalPrice: order.TotalPrice,
			Products:   orderedProduct,
		})
	}

	return &GetOrderByAccountIDResponse{
		Order: orders,
	}, nil
}

func (s *grpcServer) mustEmbedUnimplementedOrderServiceServer() {}

func (s *grpcServer) getProductDetail(ctx context.Context, productID string) (*Order_OrderProduct, error) {
	p, err := s.productClient.FindProduct(ctx, productID)
	if err != nil {
		log.Println("failed while get data product:", err)
		return nil, err
	}

	return &Order_OrderProduct{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}, nil
}
