package repository

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository/cache"
	"github.com/aWatLove/nats-lvl-zero/internal/repository/postgres"
	"gorm.io/gorm"
)

type Order interface {
	Create(order model.Order) error
	Get(uid string) (model.Order, error)
	GetAll() ([]model.Order, error)
}

type OrderCache interface {
	PutOrder(order model.Order)
	Get(uid string) (model.Order, error)
	GetAll() []model.Order
}

type Repository struct {
	Order
	OrderCache
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Order:      postgres.NewOrderPostgres(db),
		OrderCache: cache.NewOrderCache(cache.NewCache()),
	}
}
