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
	a, err := s.service.FindAccount(ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &FindAccountResponse{
		Account: a,
	}, nil
}

func (s grpcServer) GetAccounts(ctx context.Context, in *GetAccountsRequest) (*GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, in.Skip, in.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for _, a := range res {
		accounts = append(accounts, &Account{
			Id:   a.Id,
			Name: a.Name,
		})
	}
	return &GetAccountsResponse{
		Accounts: accounts,
	}, nil
}

func (s grpcServer) mustEmbedUnimplementedAccountServiceServer() {
}
