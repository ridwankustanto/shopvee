package account

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	CreateAccount(ctx context.Context, a Account) error
	GetAccountByID(ctx context.Context, id string) (*Account, error)
	ListAccounts(ctx context.Context, skip, take uint64) ([]Account, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r postgresRepository) CreateAccount(ctx context.Context, a Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts(id, name) VALUES($1, $2)", a.Id, a.Name)
	return err
}

func (r postgresRepository) GetAccountByID(ctx context.Context, id string) (*Account, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM accounts WHERE id = $1", id)
	a := &Account{}
	if err := row.Scan(&a.Id, &a.Name); err != nil {
		log.Println("failed on get account by id:", id, err)
		return nil, err
	}
	return a, nil
}

func (r postgresRepository) ListAccounts(ctx context.Context, skip, take uint64) ([]Account, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM accounts ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []Account{}
	for rows.Next() {
		a := &Account{}
		if err = rows.Scan(&a.Id, &a.Name); err == nil {
			accounts = append(accounts, *a)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}
