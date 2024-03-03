package main

import (
	"context"
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
	"os/signal"
	"sync"
	"syscall"
)

// @title WB Tech: level #0
// @version 1.0
// @description Тестовое задание. Стек: Golang, Nats-streaming, PostgreSQL

// @host localhost:8080
// @BasePath /

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

	// fill cache from db
	err = services.PutOrdersDBtoCache()
	if err != nil {
		log.Fatalf("error while loading data to cache from DB: %s", err.Error())
		return
	}

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

	go func() {
		if err = srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("server started")

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("server shutting down")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	cdb, _ := db.DB()
	if err = cdb.Close(); err != nil {
		log.Printf("error while closing db connection: %s", err.Error())
	}

	if err = sc.Close(); err != nil {
		log.Printf("error while closing nats streaming server connection: %s", err.Error())
	}

	wg.Wait()
}

// todo remove this
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
