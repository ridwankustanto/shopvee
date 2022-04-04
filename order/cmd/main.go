package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ridwankustanto/shopvee/order"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	ProductURL  string `envconfig:"PRODUCT_SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r order.Repository
	foreverSleep(3*time.Second, func(_ int) error {
		r, err = order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("failed on registering new postgres repository:", err)
			return err
		}
		return nil
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := order.NewService(r)
	log.Fatal(order.ListenGRPC(s, cfg.AccountURL, cfg.ProductURL, 8080))
}

func foreverSleep(d time.Duration, f func(int) error) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}
