package product

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	CreateProduct(ctx context.Context, product Product) error
	FindProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip, take uint64) ([]Product, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("failed on connecting to postgres server:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("failed on ping to postgres server:", err)
		return nil, err
	}

	return &postgresRepository{db: db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) CreateProduct(ctx context.Context, product Product) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO products(id, name, description, price, created_at) VALUES($1, $2, $3, $4, $5)",
		product.Id, product.Name, product.Description, product.Price, product.Timestamp)
	return err
}

func (r *postgresRepository) FindProduct(ctx context.Context, id string) (*Product, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, description, created_at, price FROM products WHERE id = $1", id)
	p := &Product{}
	if err := row.Scan(&p.Id, &p.Name, &p.Description, &p.Timestamp, &p.Price); err != nil {
		log.Println("failed on scanning product data:", err)
		return nil, err
	}
	return p, nil
}

func (r *postgresRepository) GetProducts(ctx context.Context, skip, take uint64) ([]Product, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, description, created_at, price FROM products ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		log.Println("failed on querying list product:", err)
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		p := Product{}
		if err = rows.Scan(&p.Id, &p.Name, &p.Description, &p.Timestamp, &p.Price); err == nil {
			products = append(products, p)
		}
	}
	if err = rows.Err(); err != nil {
		log.Println("failed on iterating list product:", err)
		return nil, err
	}
	return products, nil
}
