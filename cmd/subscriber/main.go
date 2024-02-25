package main

import (
	"github.com/aWatLove/nats-lvl-zero/internal/delivery/http"
	"github.com/aWatLove/nats-lvl-zero/internal/delivery/nats"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
	"github.com/aWatLove/nats-lvl-zero/internal/repository/postgres"
	"github.com/aWatLove/nats-lvl-zero/internal/service"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing .env variables^ %s", err.Error())
	}
	db, err := postgres.ConnectDB(postgres.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("failed to initialized db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	// todo fill cache from db

	// init NATS streaming server
	natsServer := nats.NewNats(services)

	// connect NATS streaming server
	sc, err := natsServer.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"), os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("error while connecting to nats streaming service: %s", err.Error())
		return
	}

	defer func(sc stan.Conn) {
		err = sc.Close()
		if err != nil {
			log.Printf("error closing nats connection: %s", err.Error())
		}
	}(sc)

	// subscribe to nats subject
	var wg sync.WaitGroup
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		err = natsServer.Subscribe(wg, sc, os.Getenv("NATS_SUBJECT"))
		if err != nil {
			return
		}
	}(&wg)
	log.Printf("successfully subscribe to nats subject")

	// init handlers
	handlers := http.NewHandler(services)
	// init server
	srv := new(http.Server)

	if err = srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}

	// todo graceful shutdown

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
