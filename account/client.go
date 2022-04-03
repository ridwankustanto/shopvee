package account

import (
	context "context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn    *grpc.ClientConn
	service AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) CreateAccount(ctx context.Context, name string) (*Account, error) {
	r, err := c.service.CreateAccount(ctx, &CreateAccountRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (c *Client) FindAccount(ctx context.Context, id string) (*Account, error) {
	a, err := c.service.FindAccount(ctx, &FindAccountRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:   a.Account.Id,
		Name: a.Account.Name,
	}, nil
}

func (c *Client) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	res, err := c.service.GetAccounts(ctx, &GetAccountsRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}
	accounts := []Account{}
	for _, a := range res.Accounts {
		accounts = append(accounts, Account{
			Id:   a.Id,
			Name: a.Name,
		})
	}
	return accounts, nil
}
