package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
	"github.com/ridwankustanto/shopvee/account"
	"github.com/ridwankustanto/shopvee/graphql/graph"
	"github.com/ridwankustanto/shopvee/graphql/graph/generated"
	"github.com/ridwankustanto/shopvee/order"
	"github.com/ridwankustanto/shopvee/product"
)

const defaultPort = "8080"

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	ProductURL string `envconfig:"PRODUCT_SERVICE_URL"`
	OrderURL   string `envconfig:"ORDER_SERVICE_URL"`
	Port       string `envconfig:"PORT"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var port string
	if cfg.Port == "" {
		port = defaultPort
	} else {
		port = cfg.Port
	}

	accountClient, err := account.NewClient(cfg.AccountURL)
	if err != nil {
		log.Fatal("failed on start new account client: ", err)
	}

	productClient, err := product.NewClient(cfg.ProductURL)
	if err != nil {
		log.Fatal("failed on start new product client: ", err)
	}

	orderClient, err := order.NewClient(cfg.OrderURL)
	if err != nil {
		log.Fatal("failed on start new order client: ", err)
	}

	s := &graph.Server{
		AccountClient: accountClient,
		ProductClient: productClient,
		OrderClient:   orderClient,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: s}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
