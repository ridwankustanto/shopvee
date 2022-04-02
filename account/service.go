package account

import (
	"context"

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
		Id:   uuid.NewString(),
		Name: name,
	}
	if err := s.repository.CreateAccount(ctx, *account); err != nil {
		return nil, err
	}
	return account, nil
}

func (a accountService) FindAccount(ctx context.Context, id string) (*Account, error) {

	return nil, nil
}

func (a accountService) GetAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {

	return nil, nil
}
