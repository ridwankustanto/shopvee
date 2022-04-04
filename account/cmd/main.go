package main

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/ridwankustanto/shopvee/account"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {

	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	foreverSleep(3*time.Second, func(_ int) error {
		r, err = account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println("failed on registering new postgres repository: ", err)
		}
		return err
	})
	defer r.Close()

	log.Println("Listening on port 8080...")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
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
