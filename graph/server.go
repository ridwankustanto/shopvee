package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
	"github.com/ridwankustanto/shopvee/account"
	"github.com/ridwankustanto/shopvee/graph/generated"
)

const defaultPort = "8080"

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
	Port       string `envconfig:"PORT"`
}

type Server struct {
	accountClient *account.Client
}

func (r *Server) Mutation() generated.MutationResolver { return &mutationResolver{server: r} }

func (r *Server) Query() generated.QueryResolver { return &queryResolver{server: r} }

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

	s := &Server{accountClient}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: s}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
