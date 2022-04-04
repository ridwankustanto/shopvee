package order

import (
	context "context"
	"database/sql"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateOrder(ctx context.Context, order Order) error
	GetOrderByAccountID(ctx context.Context, accountID string) ([]Order, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("failed on opening connection postgres:", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("failed on ping postgres server:", err)
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) CreateOrder(ctx context.Context, order Order) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("failed on declare db transaction:", err)
		return err
	}
	// Insert order
	_, err = tx.ExecContext(ctx, "INSERT INTO orders(id, account_id, total_price, timestamp) VALUES($1, $2, $3, $4)", order.Id, order.AccountId, order.TotalPrice, order.Timestamp)
	if err != nil {
		tx.Rollback()
		log.Println("failed on executing db transaction:", err)
		return err
	}

	// Insert order product
	stmt, _ := tx.PrepareContext(ctx, pq.CopyIn("order_products", "order_id", "product_id", "quantity"))
	for _, p := range order.Products {
		_, err = stmt.ExecContext(ctx, order.Id, p.Id, p.Quantity)
		if err != nil {
			tx.Rollback()
			log.Println("failed on executing db statement:", err)
			return err
		}
	}
	_, err = stmt.ExecContext(ctx)
	if err != nil {
		tx.Rollback()
		log.Println("failed on executing db statement:", err)
		return err
	}
	stmt.Close()

	return tx.Commit()
}

func (r *postgresRepository) GetOrderByAccountID(ctx context.Context, accountID string) ([]Order, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT o.id, o.account_id, o.total_price, o.timestamp, op.product_id, op.quantity
	FROM orders o JOIN order_products op ON (o.id = op.order_id)
	WHERE account_id = $1
	ORDER BY o.id`)
	if err != nil {
		log.Println("failed on quering:", err)
		return nil, err
	}
	defer rows.Close()

	mapOrder := map[string]*Order{}
	for rows.Next() {
		order := &Order{}
		orderProduct := &Order_OrderProduct{}
		if err = rows.Scan(
			&order.Id,
			&order.AccountId,
			&order.TotalPrice,
			&order.Timestamp,
			&orderProduct.Id,
			&orderProduct.Quantity,
		); err != nil {
			log.Println("failed on scan row:", err)
			return nil, err
		}

		// check if order exist on the map order
		if _, ok := mapOrder[order.Id]; ok {
			mapOrder[order.Id].Products = append(mapOrder[order.Id].Products, &Order_OrderProduct{
				Id:       orderProduct.Id,
				Quantity: orderProduct.Quantity,
			})
		} else {
			mapOrder[order.Id] = order
		}

	}

	if err = rows.Err(); err != nil {
		log.Println("failed on iterating rows:", err)
		return nil, err
	}

	orders := []Order{}
	for _, o := range mapOrder {
		orders = append(orders, Order{
			Id:         o.Id,
			AccountId:  o.AccountId,
			TotalPrice: o.TotalPrice,
			Timestamp:  o.Timestamp,
			Products:   o.Products,
		})
	}

	return orders, nil
}
