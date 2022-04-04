package test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ridwankustanto/shopvee/account"
	"github.com/ridwankustanto/shopvee/order"
	"github.com/ridwankustanto/shopvee/product"
)

var (
	accountID string
	productID string
)

func TestDBConnection(t *testing.T) {
	dbUrl := "postgres://shopvee:123456@account_db/shopvee?sslmode=disable"
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		t.Errorf("error %v", err)
		return
	}
	err = db.Ping()
	if err != nil {
		t.Errorf("error %v", err)
		return
	}
}

func TestCreateAccount(t *testing.T) {
	c, err := account.NewClient("account:8080")
	if err != nil {
		t.Errorf("Couldn't connect to the client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	a, err := c.CreateAccount(ctx, "Jonny")
	if err != nil {
		t.Errorf("Error creating account: %v", err)
		return
	}

	accountID = a.Id

	t.Logf("Success creating account: id %v & name %v", a.GetId(), a.GetName())
}

func TestFindAccount(t *testing.T) {
	c, err := account.NewClient("account:8080")
	if err != nil {
		t.Errorf("Couldn't connect to the client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	a, err := c.FindAccount(ctx, accountID)
	if err != nil {
		t.Errorf("Error finding account: %v", err)
		return
	}

	t.Logf("Success find account: id %v & name %v", a.GetId(), a.GetName())
}

func TestCreateProduct(t *testing.T) {
	c, err := product.NewClient("product:8080")
	if err != nil {
		t.Errorf("Couldn't connect to the client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p, err := c.CreateProduct(ctx, "Pisang Goreng", "Pisang goreng enak", 2000)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
		return
	}

	productID = p.Id

	t.Logf("Success creating product: id %v | name %v | description %v | price %v", p.GetId(), p.GetName(), p.GetDescription(), p.GetPrice())
}

func TestFindProduct(t *testing.T) {
	c, err := product.NewClient("product:8080")
	if err != nil {
		t.Errorf("Couldn't connect to the client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	p, err := c.FindProduct(ctx, productID)
	if err != nil {
		t.Errorf("Error finding product: %v", err)
		return
	}

	t.Logf("Success find product: id %v | name %v | description %v | price %v", p.GetId(), p.GetName(), p.GetDescription(), p.GetPrice())
}

func TestCreateOrder(t *testing.T) {
	c, err := order.NewClient("order:8080")
	if err != nil {
		t.Errorf("Couldn't connect to the client: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	products := []*order.CreateOrderRequest_OrderProduct{
		{
			ProductId: productID,
			Quantity:  10,
		},
	}

	o, err := c.CreateOrder(ctx, accountID, products)
	if err != nil {
		t.Errorf("Error creating order: %v", err)
		return
	}

	t.Logf("Success creating order: total price %v", o.TotalPrice)
}
