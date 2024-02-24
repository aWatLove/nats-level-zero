package postgres

import (
	"fmt"
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectDB(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Order{}, &model.Delivery{}, &model.Payment{}, &model.Item{})

	return db, nil
}
