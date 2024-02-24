package main

import (
	"github.com/aWatLove/nats-lvl-zero/internal/delivery/http"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
	"github.com/aWatLove/nats-lvl-zero/internal/repository/postgres"
	"github.com/aWatLove/nats-lvl-zero/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
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

	// todo init NATS

	// todo connect NATS streaming server

	// todo subscribe

	// init handlers
	handlers := http.NewHandler(services)
	// init server
	srv := new(http.Server)

	if err = srv.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
