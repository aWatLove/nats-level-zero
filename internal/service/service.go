package service

import (
	"github.com/aWatLove/nats-lvl-zero/internal/model"
	"github.com/aWatLove/nats-lvl-zero/internal/repository"
)

type Order interface {
	PutOrderDB(order model.Order) error
	PutOrderCache(order model.Order)
	GetFromDB(uid string) (model.Order, error)
	GetFromCache(uid string) (model.Order, error)
	GetAllFromDB() ([]model.Order, error)
	GetAllFromCache() ([]model.Order, error)
	PutOrdersDBtoCache() error
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
