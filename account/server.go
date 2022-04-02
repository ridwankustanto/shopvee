package account

import (
	context "context"
	"fmt"
	"net"

	"google.golang.org/grpc"
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
	// &grpcServer{s}
	RegisterAccountServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s grpcServer) CreateAccount(ctx context.Context, in *CreateAccountRequest) (*CreateAccountResponse, error) {

	a, err := s.service.CreateAccount(ctx, in.Name)
	if err != nil {
		return nil, err
	}

	return &CreateAccountResponse{
		Account: a,
	}, nil
}

func (s grpcServer) FindAccount(ctx context.Context, in *FindAccountRequest) (*FindAccountResponse, error) {

	return nil, nil
}

func (s grpcServer) GetAccounts(ctx context.Context, in *GetAccountsRequest) (*GetAccountsResponse, error) {

	return nil, nil
}

func (s grpcServer) mustEmbedUnimplementedAccountServiceServer() {
}
