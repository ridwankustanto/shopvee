package product

import (
	context "context"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	RegisterProductServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) CreateProduct(ctx context.Context, in *CreateProductRequest) (*CreateProductResponse, error) {
	p, err := s.service.CreateProduct(ctx, in.Name, in.Description, in.Price)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		Product: p,
	}, nil
}

func (s *grpcServer) FindProduct(ctx context.Context, in *FindProductRequest) (*FindProductResponse, error) {
	p, err := s.service.FindProduct(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &FindProductResponse{
		Product: p,
	}, nil
}

func (s *grpcServer) GetProducts(ctx context.Context, in *GetProductsRequest) (*GetProductsResponse, error) {
	listProduct, err := s.service.GetProducts(ctx, in.Skip, in.Take)
	if err != nil {
		return nil, err
	}

	products := []*Product{}
	for _, p := range listProduct {
		products = append(products, &Product{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Timestamp:   p.Timestamp,
		})
	}

	return &GetProductsResponse{
		Products: products,
	}, nil
}

func (s *grpcServer) mustEmbedUnimplementedProductServiceServer() {}
