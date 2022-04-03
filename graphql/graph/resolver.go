package graph

import (
	"github.com/ridwankustanto/shopvee/account"
	"github.com/ridwankustanto/shopvee/product"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Server struct {
	AccountClient *account.Client
	ProductClient *product.Client
}
