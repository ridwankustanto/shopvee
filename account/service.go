package account

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type Service interface {
	CreateAccount(ctx context.Context, name string) (*Account, error)
	FindAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error)
}

type accountService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &accountService{r}
}

func (s accountService) CreateAccount(ctx context.Context, name string) (*Account, error) {
	account := &Account{
		Id:   strings.ReplaceAll(uuid.NewString(), "-", ""),
		Name: name,
	}
	if err := s.repository.CreateAccount(ctx, *account); err != nil {
		return nil, err
	}
	return account, nil
}

func (s accountService) FindAccount(ctx context.Context, id string) (*Account, error) {
	a, err := s.repository.GetAccountByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s accountService) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	accounts, err := s.repository.ListAccounts(ctx, skip, take)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
